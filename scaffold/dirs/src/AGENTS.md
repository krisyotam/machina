# ~/src

Versioned projects. One git repo per directory. This is not a
scratch pad.

Shared primitives (hosts/SSH, subagent CLIs, models.env) live in
`~/.config/model/`. Read `~/.config/model/README.md` on launch and
append a row to `~/.local/state/model/log.csv` when the session ends.

## Put here

A project with its own history. Clone it or `git init` it. The
directory name is the repo name.

## Put somewhere else

- Notes and write-ups → `~/research`
- An experiment that is not a repo yet → `~/lab`
- Throwaway files → `~/tmp/<YYYY-MM-DD>-topic`
- Downloads → `~/dl`
- Loose scripts you want on PATH → `~/.local/bin`

## Rules

- Do not write files at the root of `~/src`. Make a repo first.
- Read the project's own `AGENTS.md` (or `CLAUDE.md`) before its
  parent. Project rules win.
- Never create a directory named `bin` or `dev` in here. Scripts
  that belong to a project live next to that project's source, not
  in a `bin/` folder.
- Never delete `~/src`.
- Do not build unless asked.
- Do not kill processes.

If you are about to start a new project and the name is not obvious,
ask. Do not invent a second top-level directory next to `src`.
