package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/krisyotam/machina/scaffold"
)

func xdg(env, home, rel string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}
	return filepath.Join(home, rel)
}

func configHome(home string) string { return xdg("XDG_CONFIG_HOME", home, ".config") }
func stateHome(home string) string  { return xdg("XDG_STATE_HOME", home, ".local/state") }
func cacheHome(home string) string  { return xdg("XDG_CACHE_HOME", home, ".cache") }

func modelConfig(home string) string   { return filepath.Join(configHome(home), "model") }
func modelState(home string) string    { return filepath.Join(stateHome(home), "model") }
func machinaConfig(home string) string { return filepath.Join(configHome(home), "machina") }
func localBin(home string) string      { return filepath.Join(home, ".local", "bin") }

type workDir struct {
	Name     string
	Purpose  string
	Protect  bool
	WipeSafe bool
	Embed    string
}

// workDirs is the table TREE.md describes. Add a row here, a template
// under scaffold/dirs/, and a line in TREE.md in the same change.
var workDirs = []workDir{
	{Name: "src", Purpose: "versioned projects, one repo each", Protect: true, Embed: "dirs/src/AGENTS.md"},
	{Name: "research", Purpose: "kept notes and investigations", Embed: "dirs/research/AGENTS.md"},
	{Name: "lab", Purpose: "experiments that are not repos yet", Embed: "dirs/lab/AGENTS.md"},
	{Name: "tmp", Purpose: "disposable work, dated subdirs", WipeSafe: true, Embed: "dirs/tmp/AGENTS.md"},
	{Name: "dl", Purpose: "downloads inbox", Protect: true, Embed: "dirs/dl/AGENTS.md"},
	{Name: "edu", Purpose: "coursework and lectures", Protect: true, Embed: "dirs/edu/AGENTS.md"},
	{Name: "lit", Purpose: "books and texts", Protect: true, Embed: "dirs/lit/AGENTS.md"},
	{Name: "sci", Purpose: "papers and datasets", Protect: true, Embed: "dirs/sci/AGENTS.md"},
	{Name: "pics", Purpose: "images", Protect: true, Embed: "dirs/pics/AGENTS.md"},
	{Name: "vids", Purpose: "video", Protect: true, Embed: "dirs/vids/AGENTS.md"},
	{Name: "music", Purpose: "audio", Protect: true, Embed: "dirs/music/AGENTS.md"},
}

var binLinks = []string{"disk", "mem", "load", "home-map"}

var agentFiles = []string{
	".claude/CLAUDE.md",
	".codex/AGENTS.md",
	".gemini/GEMINI.md",
	".grok/AGENTS.md",
	".config/opencode/AGENTS.md",
	".hermes/SOUL.md",
}

func readScaffold(name string) ([]byte, error) {
	return scaffold.FS.ReadFile(name)
}

func tilde(home, path string) string {
	if path == home {
		return "~"
	}
	sep := string(os.PathSeparator)
	if strings.HasPrefix(path, home+sep) {
		return "~/" + filepath.ToSlash(strings.TrimPrefix(path, home+sep))
	}
	return path
}

// classify reports which work directory path belongs to.
// ok is false when the path is not under any of them.
func classify(home, path string) (name, purpose string, ok bool) {
	if path == "" {
		return "", "empty path", false
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err.Error(), false
	}
	homeAbs, err := filepath.Abs(home)
	if err != nil {
		homeAbs = home
	}
	for _, d := range workDirs {
		root := filepath.Join(homeAbs, d.Name)
		if abs == root || strings.HasPrefix(abs, root+string(os.PathSeparator)) {
			return d.Name, d.Purpose, true
		}
	}
	return "", "not a machina work path", false
}

func bannedName(path string) string {
	base := strings.ToLower(filepath.Base(path))
	switch base {
	case "downloads":
		return "banned name. inbox is ~/dl"
	case "dev":
		return "banned name. code is ~/src"
	case "bin":
		return "banned name. scripts go in ~/.local/bin"
	}
	return ""
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isDir(path string) bool {
	st, err := os.Stat(path)
	return err == nil && st.IsDir()
}
