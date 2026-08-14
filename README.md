# machina

A layout for a machine that AI agents live on, plus a small Go program that
lays it down.

Agents invent their own directories, drop research next to source, fill the
disk, and ask the same questions about hosts every session. machina is a
known tree, an `AGENTS.md` in each work directory, the
[model](https://github.com/krisyotam/model) standard for shared facts, and a
few commands that print disk and memory without making anyone scrape `df -h`.

This repo is the spec. [TREE.md](TREE.md) is the tree. `docs/` explains the
pieces. `machina install` writes it onto `$HOME`.

## What it installs

1. **model.** Clones [krisyotam/model](https://github.com/krisyotam/model)
   and copies its templates into `$XDG_CONFIG_HOME/model/` and
   `$XDG_STATE_HOME/model/`. Appends the one-line pointer to each agent
   CLI's instruction file. Does not overwrite `hosts.md` or `models.env`
   if they already exist.
2. **Work directories.** Creates `src`, `research`, `lab`, `tmp`, `dl`,
   `edu`, `lit`, `sci`, `pics`, `vids`, and `music` under `$HOME` when they
   are missing, and writes a short `AGENTS.md` into each. Existing
   `AGENTS.md` files are left alone unless you pass `-f`.
3. **Check tools.** Installs `machina` into `~/.local/bin` and links
   `disk`, `mem`, `load`, and `home-map` to it. Agents run those instead
   of parsing `df` or `free`.

The machine index lives at `~/.config/machina/README.md`. Each agent CLI
gets a second one-liner pointing there, next to the model line.

## Commands

```
machina status            what is present, what is missing
machina install           model + dirs + tools
machina install model     just the model standard
machina install dirs      just work dirs and AGENTS.md
machina install bin       just ~/.local/bin
machina disk              filesystems, sizes, percent full
machina mem               memory and swap
machina load              load averages and cpu count
machina home              top-level home map
machina where PATH        which work dir a path belongs in
machina tree              print TREE.md
```

`-n` prints what would happen and writes nothing. `-home PATH` aims at a
directory other than `$HOME` (useful for tests). `-f` overwrites
`AGENTS.md` files that already exist.

## Build

```
make
make install          # PREFIX defaults to ~/.local
```

Go 1.24 or newer. No extra modules. `CGO_ENABLED=0`.

Do not build this on a Kris machine unless he said "build". `go run
./cmd/machina status` is enough to read the current box.

## What this is not

It is not a dotfile manager. srice / Stow still owns configs.

It is not a secret store. `models.env` is chmod 600 and empty until a
human fills it from Infisical or `pass`.

It is not a project template. Per-repo `AGENTS.md` files stay in the
repos. This one only covers the home layout.

It does not delete anything. Missing dirs are created. Existing files
are skipped. Protected directories (`src`, `dl`, `edu`, `lit`, `music`,
`pics`, `sci`, `vids`) are never suggested as wipe targets.

## Docs

- [docs/why.md](docs/why.md): why a machine needs this
- [docs/layout.md](docs/layout.md): what each directory is for
- [docs/model.md](docs/model.md): how the model standard is applied
- [docs/agents.md](docs/agents.md): the AGENTS.md convention
- [docs/scripts.md](docs/scripts.md): the check tools
- [docs/guards.md](docs/guards.md): hard blocks agents must not cross
- [docs/install.md](docs/install.md): install, dry-run, overwrite rules

## License

Public domain (Unlicense). It is a directory layout and a small program.
