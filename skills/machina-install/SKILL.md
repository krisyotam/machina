---
name: machina-install
description: >
  Install or repair the machina layout on a home directory. Applies
  the model standard, work dirs with AGENTS.md, check-tool links,
  and the operator skill. Use when bootstrapping a new machine,
  machina status shows missing pieces, disk/mem/home-map is not
  on PATH, or the user says install machina, scaffold, or
  bootstrap the agent home.
---

# Install the layout

Source is `github.com/krisyotam/machina`. The procedure is
`docs/install.md` in that repo. This skill is the commands.

## When to use

- A new box with no `~/.config/machina/README.md`
- `machina status` reports missing dirs, model, or bin
- `disk` / `mem` / `home-map` are not on PATH
- The human asked to install or repair machina

Do not run this because a single file is in the wrong place.
Move the file. Install is for the layout itself.

## Procedure

1. If `~/src/machina` is missing, clone it there. Creating `~/src`
   by hand is allowed; it is one of the layout names.

   ```
   git clone https://github.com/krisyotam/machina.git ~/src/machina
   cd ~/src/machina
   ```

2. Read-only pass:

   ```
   go run ./cmd/machina status
   go run ./cmd/machina install -n
   ```

3. On a Kris machine, do not `make` or `go build` unless he said
   "build". `go run` is the install path.

   ```
   go run ./cmd/machina install
   ```

   Steps, also runnable alone: `model`, `dirs`, `bin`, `skills`.

4. Fill `~/.config/model/hosts.md` and `~/.config/model/models.env`
   yourself. Real values come from Infisical or `pass`. The
   program leaves both files alone if they already exist.

5. Confirm:

   ```
   machina status
   disk
   home-map
   ```

## Overwrite

`-n` writes nothing. `-f` replaces existing `AGENTS.md` templates.
`hosts.md` and `models.env` are never overwritten.

The machine index (`~/.config/machina/README.md`) and the skill
copies under each CLI's skill dir are ours; install rewrites them.

`-home PATH` aims at a disposable tree. Use that for a dry
rehearsal, not `$HOME`.

## Pitfalls

- `go run` cannot install the `machina` binary (the temp
  executable goes away). `install bin` will say so. Links still
  work if `machina` is already on PATH; otherwise build when
  asked and re-run `install bin`.
- `MACHINA_MODEL_REPO` overrides where the model templates come
  from. Default is `~/src/model`, then a clone into cache.
- Do not run install against production paths on stargate without
  being asked. See `~/.config/model/hosts.md`.

## Verification

`machina status` shows `model`, `log`, `index`, and every work
dir as `present`. `~/.claude/skills/machina/SKILL.md` (and the
grok/codex/hermes copies) exist. `disk` prints a table.
