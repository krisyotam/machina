package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func install(w io.Writer, opt options, step string) error {
	switch step {
	case "", "all":
		if err := installModel(w, opt); err != nil {
			return err
		}
		if err := installDirs(w, opt); err != nil {
			return err
		}
		return installBin(w, opt)
	case "model":
		return installModel(w, opt)
	case "dirs":
		return installDirs(w, opt)
	case "bin":
		return installBin(w, opt)
	default:
		return fmt.Errorf("unknown install step %q (want model, dirs, bin)", step)
	}
}

func installDirs(w io.Writer, opt options) error {
	idx := filepath.Join(machinaConfig(opt.home), "README.md")
	body, err := readScaffold("machine/README.md")
	if err != nil {
		return err
	}
	if err := writeFile(w, opt, idx, body, 0o644, true); err != nil {
		return err
	}

	for _, d := range workDirs {
		dir := filepath.Join(opt.home, d.Name)
		if err := ensureDir(w, opt, dir); err != nil {
			return err
		}
		body, err := readScaffold(d.Embed)
		if err != nil {
			return fmt.Errorf("%s: %w", d.Embed, err)
		}
		if err := writeFile(w, opt, filepath.Join(dir, "AGENTS.md"), body, 0o644, opt.force); err != nil {
			return err
		}
	}

	line := machinaLine(opt.home)
	needle := machinaConfig(opt.home)
	if strings.HasPrefix(needle, opt.home) {
		needle = tilde(opt.home, needle)
	}
	return appendAgentLine(w, opt, needle+"/", line)
}

func installBin(w io.Writer, opt options) error {
	binDir := localBin(opt.home)
	if err := ensureDir(w, opt, binDir); err != nil {
		return err
	}

	dest := filepath.Join(binDir, "machina")
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("executable: %w", err)
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		exe, _ = os.Executable()
	}

	if fromGoRun(exe) {
		printAct(w, "skip", dest, "running from go run; build and re-run install bin")
	} else if err := copyExe(w, opt, exe, dest); err != nil {
		return err
	}

	for _, name := range binLinks {
		link := filepath.Join(binDir, name)
		if err := ensureLink(w, opt, "machina", link); err != nil {
			return err
		}
	}
	return nil
}

func status(w io.Writer, opt options) error {
	type row struct{ k, v, extra string }
	var rows []row

	mc := modelConfig(opt.home)
	if isDir(mc) {
		rows = append(rows, row{"model", "present", tilde(opt.home, mc)})
	} else {
		rows = append(rows, row{"model", "missing", tilde(opt.home, mc)})
	}
	logp := filepath.Join(modelState(opt.home), "log.csv")
	if exists(logp) {
		rows = append(rows, row{"log", "present", tilde(opt.home, logp)})
	} else {
		rows = append(rows, row{"log", "missing", tilde(opt.home, logp)})
	}
	idx := filepath.Join(machinaConfig(opt.home), "README.md")
	if exists(idx) {
		rows = append(rows, row{"index", "present", tilde(opt.home, idx)})
	} else {
		rows = append(rows, row{"index", "missing", tilde(opt.home, idx)})
	}
	bin := filepath.Join(localBin(opt.home), "machina")
	if exists(bin) {
		rows = append(rows, row{"bin", "present", tilde(opt.home, bin)})
	} else {
		rows = append(rows, row{"bin", "missing", tilde(opt.home, bin)})
	}

	for _, d := range workDirs {
		dir := filepath.Join(opt.home, d.Name)
		agents := filepath.Join(dir, "AGENTS.md")
		switch {
		case !isDir(dir):
			rows = append(rows, row{d.Name, "missing", ""})
		case exists(agents):
			rows = append(rows, row{d.Name, "present", "AGENTS.md"})
		default:
			rows = append(rows, row{d.Name, "present", "no AGENTS.md"})
		}
	}

	kw, vw := 0, 0
	for _, r := range rows {
		if len(r.k) > kw {
			kw = len(r.k)
		}
		if len(r.v) > vw {
			vw = len(r.v)
		}
	}
	for _, r := range rows {
		if r.extra == "" {
			fmt.Fprintf(w, "%-*s  %-*s\n", kw, r.k, vw, r.v)
		} else {
			fmt.Fprintf(w, "%-*s  %-*s  %s\n", kw, r.k, vw, r.v, r.extra)
		}
	}
	return nil
}

func ensureDir(w io.Writer, opt options, dir string) error {
	if isDir(dir) {
		printAct(w, "ok", tilde(opt.home, dir), "exists")
		return nil
	}
	if opt.dry {
		printAct(w, "mkdir", tilde(opt.home, dir), "")
		return nil
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	printAct(w, "mkdir", tilde(opt.home, dir), "")
	return nil
}

// writeFile writes body to path. If overwrite is false and the file
// exists, it is left alone. The machine index is the one file we
// rewrite on purpose (overwrite=true).
func writeFile(w io.Writer, opt options, path string, body []byte, mode os.FileMode, overwrite bool) error {
	if exists(path) && !overwrite {
		printAct(w, "skip", tilde(opt.home, path), "exists")
		return nil
	}
	if opt.dry {
		verb := "write"
		if exists(path) {
			verb = "replace"
		}
		printAct(w, verb, tilde(opt.home, path), "")
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(path, body, mode); err != nil {
		return err
	}
	printAct(w, "write", tilde(opt.home, path), "")
	return nil
}

func copyExe(w io.Writer, opt options, src, dest string) error {
	if sameFile(src, dest) {
		printAct(w, "ok", tilde(opt.home, dest), "already this binary")
		return nil
	}
	if opt.dry {
		printAct(w, "copy", tilde(opt.home, dest), "")
		return nil
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	tmp := dest + ".tmp"
	out, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(out, in)
	closeErr := out.Close()
	if copyErr != nil {
		os.Remove(tmp)
		return copyErr
	}
	if closeErr != nil {
		os.Remove(tmp)
		return closeErr
	}
	if err := os.Rename(tmp, dest); err != nil {
		os.Remove(tmp)
		return err
	}
	printAct(w, "copy", tilde(opt.home, dest), "")
	return nil
}

func ensureLink(w io.Writer, opt options, target, link string) error {
	disp := tilde(opt.home, link)
	if st, err := os.Lstat(link); err == nil {
		if st.Mode()&os.ModeSymlink != 0 {
			cur, _ := os.Readlink(link)
			if cur == target {
				printAct(w, "ok", disp, "-> "+target)
				return nil
			}
			// Only replace a link we previously made (name is one of ours
			// and currently points at machina or an old relative name).
			if cur == "machina" || filepath.Base(cur) == "machina" {
				if opt.dry {
					printAct(w, "relink", disp, "-> "+target)
					return nil
				}
				if err := os.Remove(link); err != nil {
					return err
				}
			} else {
				printAct(w, "skip", disp, "symlink points at "+cur)
				return nil
			}
		} else {
			printAct(w, "skip", disp, "not a symlink")
			return nil
		}
	}
	if opt.dry {
		printAct(w, "link", disp, "-> "+target)
		return nil
	}
	if err := os.Symlink(target, link); err != nil {
		return err
	}
	printAct(w, "link", disp, "-> "+target)
	return nil
}

func printAct(w io.Writer, verb, path, note string) {
	if note == "" {
		fmt.Fprintf(w, "%-8s %s\n", verb, path)
		return
	}
	fmt.Fprintf(w, "%-8s %s (%s)\n", verb, path, note)
}

func fromGoRun(exe string) bool {
	if strings.Contains(exe, string(os.PathSeparator)+"go-build") {
		return true
	}
	tmp := os.TempDir()
	return tmp != "" && strings.HasPrefix(exe, tmp)
}

func sameFile(a, b string) bool {
	sa, err := os.Stat(a)
	if err != nil {
		return false
	}
	sb, err := os.Stat(b)
	if err != nil {
		return false
	}
	return os.SameFile(sa, sb)
}

func machinaLine(home string) string {
	cfg := tilde(home, machinaConfig(home))
	return "Machine layout and check tools live in `" + cfg + "/`. Read `" + cfg + "/README.md` on launch. Run `disk`, `mem`, `load`, `home-map` from ~/.local/bin instead of scraping df/free."
}
