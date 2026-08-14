package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

var skipFS = map[string]bool{
	"autofs": true, "binfmt_misc": true, "bpf": true, "cgroup": true,
	"cgroup2": true, "configfs": true, "debugfs": true, "devpts": true,
	"devtmpfs": true, "efivarfs": true, "fusectl": true, "hugetlbfs": true,
	"mqueue": true, "nsfs": true, "overlay": true, "proc": true,
	"pstore": true, "ramfs": true, "rpc_pipefs": true, "securityfs": true,
	"squashfs": true, "sysfs": true, "tmpfs": true, "tracefs": true,
}

func reportDisk(w io.Writer, all bool) error {
	mounts, err := readMounts()
	if err != nil {
		return err
	}

	type row struct {
		dev, fstype, size, used, avail, pct, mount string
		full                                       bool
	}
	var rows []row
	seen := map[string]bool{}
	for _, m := range mounts {
		if !all {
			if skipFS[m.fstype] {
				continue
			}
			if skipMount(m.dir) {
				continue
			}
		}
		key := m.dev + "\t" + m.dir
		if seen[key] {
			continue
		}
		seen[key] = true

		var st syscall.Statfs_t
		if err := syscall.Statfs(m.dir, &st); err != nil {
			continue
		}
		bsize := uint64(st.Bsize)
		total := st.Blocks * bsize
		avail := st.Bavail * bsize
		free := st.Bfree * bsize
		used := total - free
		pctn := 0
		if used+avail > 0 {
			pctn = int((used*100 + (used+avail)/2) / (used + avail))
		}
		rows = append(rows, row{
			dev:    m.dev,
			fstype: m.fstype,
			size:   fmtBytes(total),
			used:   fmtBytes(used),
			avail:  fmtBytes(avail),
			pct:    fmt.Sprintf("%d%%", pctn),
			mount:  m.dir,
			full:   pctn >= 90,
		})
	}

	cols := []struct {
		name string
		get  func(row) string
	}{
		{"device", func(r row) string { return r.dev }},
		{"type", func(r row) string { return r.fstype }},
		{"size", func(r row) string { return r.size }},
		{"used", func(r row) string { return r.used }},
		{"avail", func(r row) string { return r.avail }},
		{"use%", func(r row) string { return r.pct }},
		{"mount", func(r row) string { return r.mount }},
	}
	widths := make([]int, len(cols))
	for i, c := range cols {
		widths[i] = len(c.name)
	}
	for _, r := range rows {
		for i, c := range cols {
			if n := len(c.get(r)); n > widths[i] {
				widths[i] = n
			}
		}
	}
	for i, c := range cols {
		if i > 0 {
			fmt.Fprint(w, "  ")
		}
		fmt.Fprintf(w, "%-*s", widths[i], c.name)
	}
	fmt.Fprintln(w)
	var warns []string
	for _, r := range rows {
		for i, c := range cols {
			if i > 0 {
				fmt.Fprint(w, "  ")
			}
			fmt.Fprintf(w, "%-*s", widths[i], c.get(r))
		}
		fmt.Fprintln(w)
		if r.full {
			warns = append(warns, r.mount+" "+r.pct)
		}
	}
	if len(warns) == 0 {
		fmt.Fprintln(w, "warn  none")
	} else {
		fmt.Fprintln(w, "warn  "+strings.Join(warns, ", "))
	}
	return nil
}

type mount struct {
	dev, dir, fstype string
}

func readMounts() ([]mount, error) {
	f, err := os.Open("/proc/self/mounts")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []mount
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) < 3 {
			continue
		}
		out = append(out, mount{
			dev:    unescapeMount(fields[0]),
			dir:    unescapeMount(fields[1]),
			fstype: fields[2],
		})
	}
	return out, sc.Err()
}

func skipMount(dir string) bool {
	for _, p := range []string{"/proc", "/sys", "/dev", "/run"} {
		if dir == p || strings.HasPrefix(dir, p+"/") {
			return true
		}
	}
	return false
}

func unescapeMount(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+3 < len(s) && isOctal(s[i+1]) && isOctal(s[i+2]) && isOctal(s[i+3]) {
			n := int(s[i+1]-'0')*64 + int(s[i+2]-'0')*8 + int(s[i+3]-'0')
			b.WriteByte(byte(n))
			i += 3
			continue
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

func isOctal(c byte) bool { return c >= '0' && c <= '7' }

func reportMem(w io.Writer) error {
	info, err := readMeminfo()
	if err != nil {
		return err
	}
	total := info["MemTotal"]
	avail := info["MemAvailable"]
	if avail == 0 {
		avail = info["MemFree"]
	}
	used := uint64(0)
	if total > avail {
		used = total - avail
	}
	swapT := info["SwapTotal"]
	swapF := info["SwapFree"]
	swapU := uint64(0)
	if swapT > swapF {
		swapU = swapT - swapF
	}
	fmt.Fprintf(w, "total  %s\n", fmtBytes(total))
	fmt.Fprintf(w, "used   %s\n", fmtBytes(used))
	fmt.Fprintf(w, "avail  %s\n", fmtBytes(avail))
	fmt.Fprintf(w, "swap   %s / %s\n", fmtBytes(swapU), fmtBytes(swapT))
	return nil
}

func readMeminfo() (map[string]uint64, error) {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	out := map[string]uint64{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		// MemTotal:       12345 kB
		k, rest, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		fields := strings.Fields(rest)
		if len(fields) == 0 {
			continue
		}
		n, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			continue
		}
		if len(fields) > 1 && strings.EqualFold(fields[1], "kB") {
			n *= 1024
		}
		out[k] = n
	}
	return out, sc.Err()
}

func reportLoad(w io.Writer) error {
	b, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return err
	}
	fields := strings.Fields(string(b))
	if len(fields) < 3 {
		return fmt.Errorf("short /proc/loadavg")
	}
	fmt.Fprintf(w, "load   %s %s %s\n", fields[0], fields[1], fields[2])
	fmt.Fprintf(w, "cpus   %d\n", runtime.NumCPU())
	return nil
}

func reportHome(w io.Writer, opt options) error {
	type row struct{ dir, purpose, agents, size string }
	var rows []row
	for _, d := range workDirs {
		p := filepath.Join(opt.home, d.Name)
		r := row{dir: "~/" + d.Name, purpose: d.Purpose, agents: "-", size: "-"}
		if !isDir(p) {
			r.agents = "missing"
			rows = append(rows, r)
			continue
		}
		if exists(filepath.Join(p, "AGENTS.md")) {
			r.agents = "yes"
		} else {
			r.agents = "no"
		}
		if opt.sizes {
			n, err := dirSize(p)
			if err != nil {
				r.size = "?"
			} else {
				r.size = fmtBytes(n)
			}
		}
		rows = append(rows, r)
	}

	dw, pw, aw, sw := len("dir"), len("purpose"), len("agents"), len("size")
	for _, r := range rows {
		if len(r.dir) > dw {
			dw = len(r.dir)
		}
		if len(r.purpose) > pw {
			pw = len(r.purpose)
		}
		if len(r.agents) > aw {
			aw = len(r.agents)
		}
		if len(r.size) > sw {
			sw = len(r.size)
		}
	}
	if opt.sizes {
		fmt.Fprintf(w, "%-*s  %-*s  %-*s  %-*s\n", dw, "dir", pw, "purpose", aw, "agents", sw, "size")
		for _, r := range rows {
			fmt.Fprintf(w, "%-*s  %-*s  %-*s  %-*s\n", dw, r.dir, pw, r.purpose, aw, r.agents, sw, r.size)
		}
		return nil
	}
	fmt.Fprintf(w, "%-*s  %-*s  %-*s\n", dw, "dir", pw, "purpose", aw, "agents")
	for _, r := range rows {
		fmt.Fprintf(w, "%-*s  %-*s  %-*s\n", dw, r.dir, pw, r.purpose, aw, r.agents)
	}
	return nil
}

func reportWhere(w io.Writer, home, path string) {
	if msg := bannedName(path); msg != "" {
		fmt.Fprintf(w, "path   %s\n", path)
		fmt.Fprintf(w, "dir    -\n")
		fmt.Fprintf(w, "note   %s\n", msg)
		return
	}
	name, purpose, ok := classify(home, path)
	fmt.Fprintf(w, "path   %s\n", path)
	if !ok {
		fmt.Fprintf(w, "dir    -\n")
		fmt.Fprintf(w, "note   %s\n", purpose)
		return
	}
	fmt.Fprintf(w, "dir    ~/%s\n", name)
	fmt.Fprintf(w, "note   %s\n", purpose)
}

func dirSize(root string) (uint64, error) {
	var n uint64
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.Mode().IsRegular() {
			n += uint64(info.Size())
		}
		return nil
	})
	return n, err
}

func fmtBytes(n uint64) string {
	if n < 1024 {
		return fmt.Sprintf("%dB", n)
	}
	units := []string{"K", "M", "G", "T", "P"}
	f := float64(n)
	i := 0
	for f >= 1024 && i < len(units)-1 {
		f /= 1024
		i++
	}
	if f >= 10 || i == 0 {
		return fmt.Sprintf("%.0f%s", f, units[i])
	}
	return fmt.Sprintf("%.1f%s", f, units[i])
}
