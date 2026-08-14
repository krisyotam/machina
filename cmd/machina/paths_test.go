package main

import (
	"path/filepath"
	"testing"
)

func TestClassify(t *testing.T) {
	home := "/home/kris"
	cases := []struct {
		path, name string
		ok         bool
	}{
		{home + "/src", "src", true},
		{home + "/src/machina/README.md", "src", true},
		{home + "/tmp/2026-08-14-x", "tmp", true},
		{home + "/research/notes/a.md", "research", true},
		{home + "/lab", "lab", true},
		{home + "/.config/model", "", false},
		{home, "", false},
		{"/etc/passwd", "", false},
	}
	for _, c := range cases {
		name, _, ok := classify(home, c.path)
		if ok != c.ok || name != c.name {
			t.Errorf("classify(%q) = %q, %v; want %q, %v", c.path, name, ok, c.name, c.ok)
		}
	}
}

func TestBannedName(t *testing.T) {
	if got := bannedName("/home/kris/Downloads"); got == "" {
		t.Fatal("Downloads should be banned")
	}
	if got := bannedName(filepath.Join("foo", "dev")); got == "" {
		t.Fatal("dev should be banned")
	}
	if got := bannedName("/home/kris/src"); got != "" {
		t.Fatalf("src should not be banned, got %q", got)
	}
}

func TestTilde(t *testing.T) {
	home := "/home/kris"
	if got := tilde(home, home); got != "~" {
		t.Fatalf("home itself: %q", got)
	}
	if got := tilde(home, home+"/src/machina"); got != "~/src/machina" {
		t.Fatalf("nested: %q", got)
	}
	if got := tilde(home, "/etc"); got != "/etc" {
		t.Fatalf("outside: %q", got)
	}
}

func TestUnescapeMount(t *testing.T) {
	got := unescapeMount(`/mnt/some\040dir`)
	if got != "/mnt/some dir" {
		t.Fatalf("got %q", got)
	}
}

func TestEmbedPaths(t *testing.T) {
	if _, err := readScaffold("TREE.md"); err != nil {
		t.Fatalf("TREE.md: %v", err)
	}
	if _, err := readScaffold("machine/README.md"); err != nil {
		t.Fatalf("machine/README.md: %v", err)
	}
	for _, d := range workDirs {
		if _, err := readScaffold(d.Embed); err != nil {
			t.Errorf("%s: %v", d.Embed, err)
		}
	}
	for _, name := range []string{"machina", "machina-install"} {
		if _, err := readSkill(name); err != nil {
			t.Errorf("skill %s: %v", name, err)
		}
	}
}

func TestFmtBytes(t *testing.T) {
	if got := fmtBytes(500); got != "500B" {
		t.Fatalf("500: %s", got)
	}
	if got := fmtBytes(1024); got != "1.0K" {
		t.Fatalf("1024: %s", got)
	}
	if got := fmtBytes(10 * 1024 * 1024); got != "10M" {
		t.Fatalf("10M: %s", got)
	}
}
