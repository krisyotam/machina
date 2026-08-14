# Install

From a checkout:

```
go run ./cmd/machina status          # read-only report
go run ./cmd/machina install -n      # dry-run
go run ./cmd/machina install         # write
```

After `make install` (PREFIX defaults to `~/.local`):

```
machina status
machina install
```

On a Kris machine, do not run `make` unless he said "build".

## Order

`machina install` with no extra word does four steps:

1. `model`: clone/apply [krisyotam/model](https://github.com/krisyotam/model)
2. `dirs`: create work directories and write missing `AGENTS.md`
3. `bin`: install `machina` into `~/.local/bin` and link the short names
4. `skills`: copy the operator and install skills into each CLI

Each step can be run alone: `machina install model`,
`machina install dirs`, `machina install bin`,
`machina install skills`.

## Flags

| Flag | Meaning |
|------|---------|
| `-n` | Print what would happen. Write nothing. |
| `-f` | Overwrite existing `AGENTS.md` templates. Still never overwrites `hosts.md` or `models.env`. |
| `-home PATH` | Treat PATH as `$HOME`. Used in tests. |
| `-a` | `disk` only: include virtual filesystems. |
| `-s` | `home-map` only: compute directory sizes. |

## Overwrite rules

| Path | Missing | Exists |
|------|---------|--------|
| `~/src` and the other work dirs | created | left alone |
| `~/src/AGENTS.md` (etc.) | written from template | left alone, unless `-f` |
| `~/.config/machina/README.md` | written | rewritten (this file is ours) |
| `~/.config/model/README.md` | copied from model | left alone |
| `~/.config/model/hosts.md` | copied from template | never overwritten |
| `~/.config/model/models.env` | copied from example, mode 0600 | never overwritten |
| `~/.local/state/model/log.csv` | copied header | left alone |
| `~/.local/bin/machina` | installed | replaced with the current binary when `install bin` can see itself |
| `~/.local/bin/disk` (and friends) | symlinked to `machina` | replaced if they already point at machina; left alone otherwise |
| `~/.claude/skills/machina/SKILL.md` (and the other CLI copies) | written | rewritten (these files are ours) |

`install bin` copies the running executable when it can. If you are
on `go run`, it prints that you need a built binary for the `bin`
step and continues with the links only if `machina` is already on
PATH.

## Aiming at another home

```
go run ./cmd/machina -home /tmp/fake-home install -n
```

Nothing outside that path is touched. Useful for reading the
program against a disposable tree.

## First boot on a new machine

1. Clone this repo into `~/src/machina` (create `~/src` by hand if
   you have to).
2. `go run ./cmd/machina install`
3. Edit `~/.config/model/hosts.md` and `~/.config/model/models.env`.
4. Confirm `disk` and `home-map` print something sane, and that
   `~/.claude/skills/machina/SKILL.md` exists.

The work directories will be empty except for `AGENTS.md`. That is
the point: the names exist before an agent has to invent them.
