package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const modelRepoURL = "https://github.com/krisyotam/model.git"

// Files copied from the model template into XDG config. hosts.md and
// models.env are never overwritten once they exist.
var modelConfigFiles = []struct {
	src, dst string
	mode     os.FileMode
	protect  bool
}{
	{src: "template/README.md", dst: "README.md", mode: 0o644},
	{src: "template/hosts.md", dst: "hosts.md", mode: 0o644, protect: true},
	{src: "template/subagents.md", dst: "subagents.md", mode: 0o644},
	{src: "template/models.md", dst: "models.md", mode: 0o644},
	{src: "template/models.env.example", dst: "models.env", mode: 0o600, protect: true},
}

func installModel(w io.Writer, opt options) error {
	src, err := findModelRepo(w, opt)
	if err != nil {
		return err
	}

	cfg := modelConfig(opt.home)
	st := modelState(opt.home)
	if err := ensureDir(w, opt, cfg); err != nil {
		return err
	}
	if err := ensureDir(w, opt, st); err != nil {
		return err
	}

	if opt.dry && !looksLikeModel(src) {
		for _, f := range modelConfigFiles {
			printAct(w, "write", tilde(opt.home, filepath.Join(cfg, f.dst)), "from model template")
		}
		printAct(w, "write", tilde(opt.home, filepath.Join(st, "log.csv")), "from model template")
		line := modelLine(opt.home)
		needle := tilde(opt.home, cfg) + "/"
		return appendAgentLine(w, opt, needle, line)
	}

	for _, f := range modelConfigFiles {
		body, err := os.ReadFile(filepath.Join(src, f.src))
		if err != nil {
			return fmt.Errorf("model template %s: %w", f.src, err)
		}
		dst := filepath.Join(cfg, f.dst)
		overwrite := !f.protect && opt.force
		if err := writeFile(w, opt, dst, body, f.mode, overwrite); err != nil {
			return err
		}
		if !opt.dry && f.mode == 0o600 && exists(dst) {
			_ = os.Chmod(dst, 0o600)
		}
	}

	logBody, err := os.ReadFile(filepath.Join(src, "template/log.csv"))
	if err != nil {
		return fmt.Errorf("model template log.csv: %w", err)
	}
	if err := writeFile(w, opt, filepath.Join(st, "log.csv"), logBody, 0o644, false); err != nil {
		return err
	}

	line := modelLine(opt.home)
	needle := modelConfig(opt.home)
	if strings.HasPrefix(needle, opt.home) {
		needle = tilde(opt.home, needle)
	}
	return appendAgentLine(w, opt, needle+"/", line)
}

func findModelRepo(w io.Writer, opt options) (string, error) {
	if p := os.Getenv("MACHINA_MODEL_REPO"); p != "" {
		if looksLikeModel(p) {
			printAct(w, "ok", p, "MACHINA_MODEL_REPO")
			return p, nil
		}
		return "", fmt.Errorf("MACHINA_MODEL_REPO=%s does not look like krisyotam/model", p)
	}
	local := filepath.Join(opt.home, "src", "model")
	if looksLikeModel(local) {
		printAct(w, "ok", tilde(opt.home, local), "existing checkout")
		return local, nil
	}
	cache := filepath.Join(cacheHome(opt.home), "machina", "model")
	if looksLikeModel(cache) {
		printAct(w, "ok", tilde(opt.home, cache), "cached clone")
		return cache, nil
	}
	if exists(cache) {
		return "", fmt.Errorf("%s exists and is not a model checkout; remove it by hand", cache)
	}
	if opt.dry {
		printAct(w, "clone", modelRepoURL, "-> "+tilde(opt.home, cache))
		return cache, nil
	}
	if err := os.MkdirAll(filepath.Dir(cache), 0o755); err != nil {
		return "", err
	}
	cmd := exec.Command("git", "clone", "--depth", "1", modelRepoURL, cache)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git clone %s: %w", modelRepoURL, err)
	}
	printAct(w, "clone", tilde(opt.home, cache), "")
	return cache, nil
}

func looksLikeModel(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "template", "README.md"))
	return err == nil
}

func modelLine(home string) string {
	cfg := tilde(home, modelConfig(home))
	st := tilde(home, modelState(home))
	return "Shared primitives (hosts/SSH, subagent CLIs, models.env) live in `" + cfg + "/`. Read `" + cfg + "/README.md` on launch and append a row to `" + st + "/log.csv` when the session ends."
}

func appendAgentLine(w io.Writer, opt options, needle, line string) error {
	for _, rel := range agentFiles {
		path := filepath.Join(opt.home, rel)
		if err := appendLine(w, opt, path, needle, line); err != nil {
			return err
		}
	}
	return nil
}

func appendLine(w io.Writer, opt options, path, needle, line string) error {
	disp := tilde(opt.home, path)
	if exists(path) {
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(b), needle) {
			printAct(w, "ok", disp, "already points here")
			return nil
		}
	}
	if opt.dry {
		printAct(w, "append", disp, "")
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	// Blank line before the pointer if the file already has content
	// that does not end in a newline.
	if st, err := f.Stat(); err == nil && st.Size() > 0 {
		if _, err := f.WriteString("\n"); err != nil {
			return err
		}
	}
	if _, err := f.WriteString(line + "\n"); err != nil {
		return err
	}
	printAct(w, "append", disp, "")
	return nil
}
