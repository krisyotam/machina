# AGENTS.md

How to work on machina. Shared machine facts live in
`~/.config/model/`. Read `~/.config/model/README.md` on launch and
append a row to `~/.local/state/model/log.csv` when the session ends.

## What this repo is

A directory layout plus a Go program that writes that layout onto
`$HOME`. [TREE.md](TREE.md) is the spec. `docs/` explains it.
`scaffold/` is the text that gets written. `cmd/machina` is the
program.

## Rules

- TREE.md and `scaffold/dirs/` stay in sync. If you add a work
  directory, add the row and the `AGENTS.md` template, then the
  entry in `cmd/machina/paths.go` (`workDirs`), all in the same
  change.
- Do not add a config format, a daemon, or a plugin system. If a
  primitive needs that, it does not belong here.
- Facts about hosts and other agent CLIs stay in
  [krisyotam/model](https://github.com/krisyotam/model). machina
  installs that repo. It does not fork its contents.
- Never create a directory named `bin` or `dev` in this tree. Templates
  for `~/.local/bin` live under `scaffold/` only as documentation;
  the program writes the links itself.
- Do not run `make`, `go build`, or `go test` on a Kris machine unless
  he said "build". `gofmt` is fine.
- Do not delete `~/src`, `~/dl`, `~/edu`, `~/lit`, `~/music`, `~/pics`,
  `~/sci`, or `~/vids`. Do not create `~/Downloads` or `~/dev`.
- Do not kill processes. Do not start remaps or remapd.

## Layout

```
TREE.md                 symlink to scaffold/TREE.md (the machine tree)
docs/                   why, layout, model, agents, scripts, guards, install
scaffold/
  fs.go                 go:embed of the templates
  machine/README.md     becomes ~/.config/machina/README.md
  dirs/<name>/AGENTS.md becomes ~/<name>/AGENTS.md
cmd/machina/
  main.go               flags and dispatch
  paths.go              XDG paths, work dir table, classify
  install.go            write dirs, AGENTS.md, bin links
  model.go              clone and apply krisyotam/model
  report.go             disk, mem, load, home
```

## Adding a work directory

1. Add a short `AGENTS.md` at `scaffold/dirs/<name>/AGENTS.md`.
2. Add the row to TREE.md and docs/layout.md.
3. Add a `workDir` entry in `cmd/machina/paths.go`.
4. Say whether it is protected (never a wipe target).

## Adding a check command

Add a case in `cmd/machina/main.go` and a function in `report.go`.
If agents should run it by a short name, add that name to `binLinks`
in `paths.go`. Output must be a small table or `key  value` lines.
No banners, no color, no unicode meters.

## Style

Stdlib only. `flag`, not cobra. Log prefix `machina: `. Dry-run is
`-n`. The program must not call `RemoveAll` or rewrite
`~/.config/model/hosts.md` / `models.env`.
