# AGENTS.md

machina treats `AGENTS.md` the way Claude Code treats `CLAUDE.md`:
the file in the directory you just entered is the law for that
directory. Sophia's repos already use this name. So do Codex and
several other CLIs.

## Two layers

1. **Machine.** `~/.config/machina/README.md` is the index of the
   home layout. Each agent CLI gets one line pointing at it, next
   to the model line.
2. **Directory.** `~/src/AGENTS.md`, `~/research/AGENTS.md`, and the
   rest say what that directory is for. A project's own `AGENTS.md`
   (inside a repo) still wins for work inside that repo.

Read in this order: the project's file, then the work-directory
file, then the machine index, then `~/.config/model/README.md`.
Closer wins.

## What a directory file contains

Short. An agent will read it every time it `cd`s in.

- What belongs here
- What belongs somewhere else
- Naming (dated tmp dirs, one repo per src dir, one topic per
  research dir)
- Cleanup (wipe-safe or protected)
- The pointer line to `~/.config/model/`

No voice, no personality, no essay. Those belong in a project's
`SOUL.md` if the project has one.

## Recognition line

`machina install` appends this to the same agent instruction files
that model uses (`~/.claude/CLAUDE.md`, `~/.codex/AGENTS.md`,
`~/.gemini/GEMINI.md`, `~/.grok/AGENTS.md`,
`~/.config/opencode/AGENTS.md`, `~/.hermes/SOUL.md`):

```
Machine layout and check tools live in `~/.config/machina/`. Read `~/.config/machina/README.md` on launch. Run `disk`, `mem`, `load`, `home-map` from ~/.local/bin instead of scraping df/free.
```

If the file already contains `~/.config/machina/`, the line is not
added again.

Skills (`skills/machina`, `skills/machina-install`) are the
procedure a model loads when it has not opened this repo. See
[skills.md](skills.md).

## Editing a live AGENTS.md

The copies under `$HOME` are the live files. The copies under
`scaffold/dirs/` are the templates. Change the template, then
`machina install dirs -f` to push it out. Without `-f`, an existing
file is left alone so local notes survive.
