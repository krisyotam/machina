---
name: machina
description: >
  Operate a machina-managed machine. Read the home layout, put work
  in the right directory, inspect disk/memory/load with the check
  tools, and respect the hard blocks. Use when landing on a box,
  deciding where a file goes, checking disk or memory, starting a
  session on a Kris machine, or the user mentions machina, home
  layout, work dirs, or AGENTS.md directories.
---

# Operate the machine

machina is the home layout. Facts (hosts, other CLIs, model slugs)
are the sibling standard in `~/.config/model/`. Do not invent a
third place for either.

## When to use

- Session start on a box that might have this layout
- Deciding where a file, clone, note, or download goes
- Checking disk, memory, or load
- The user says machina, home layout, or "where does this go"

## Session start

1. Read `~/.config/model/README.md`.
2. Read `~/.config/machina/README.md`. If that file is missing, run
   `machina status` and see `machina-install`.
3. Do not copy facts out of those files into the project.

## Where work goes

Run `home-map` or `machina where PATH`. Then read
`~/<dir>/AGENTS.md` before writing into that directory.

The live table is `~/.config/machina/README.md`. Do not invent a
new top-level name. Do not write at the root of `$HOME`.

| Kind | Dest |
|------|------|
| git repo | `~/src/<name>` |
| kept notes | `~/research/<topic>` |
| experiment, not a repo yet | `~/lab/<topic>` |
| throwaway | `~/tmp/<YYYY-MM-DD>-<topic>` |
| unsorted download | `~/dl` (sort out this session) |

A project's own `AGENTS.md` wins inside that repo.

## Check tools

Do not scrape `df`, `free`, or `uptime`. From `~/.local/bin`:

```
disk          filesystems; warn line at 90% full
mem           total / used / avail / swap
load          1 5 15 minute load, cpu count
home-map      work dirs and whether AGENTS.md is present
machina where PATH
```

`disk` before writing anything large. `home-map -s` walks trees
and can take minutes; skip `-s` unless you need sizes.

If the commands are missing, the layout is not installed. See
`machina-install`.

## Guards

Hard blocks. Stop and say so.

- Do not delete `~/src`, `~/dl`, `~/edu`, `~/lit`, `~/music`,
  `~/pics`, `~/sci`, `~/vids`.
- Do not create `bin/` anywhere, `~/dev`, or `Downloads/`.
  Scripts go in `~/.local/bin`. Code goes in `~/src`.
- Do not run `remaps` or `remapd`.
- Do not build (`make`, `cargo build`, `npm run build`) unless
  the human said "build".
- Do not kill a process unless it was named in this conversation.
- Do not create a new top-level name under `$HOME`.

Long form: `docs/guards.md` in the machina repo, or the Guards
section of `~/.config/machina/README.md`.

## Session end

Append one row to `~/.local/state/model/log.csv`. Format and slug
rules are in `~/.config/model/README.md`. Mint a missing slug in
`models.md` before writing the row.

## Pitfalls

- `~/lab` is not `~/research`. Promote or delete; do not let lab
  become an unnamed junk drawer.
- `~/tmp` uses dated subdirs. A dir still there a week later
  should have been promoted or removed.
- Site images for krisyotam.com do not go in `~/pics`. CDN rules
  live in `~/.claude/docs/cdn.md`.

## Verification

```
machina status
disk
home-map
```

`status` should show model, index, and the work dirs as present.
`disk` should end with `warn  none` or a mount you already knew
about.
