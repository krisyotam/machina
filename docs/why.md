# Why

An agent CLI is a stranger every time it starts. It does not remember
that research lives in `~/research`, that `~/tmp` is wipe-safe, or that
`df -h` on this box is noisy. It will invent a `projects/` directory,
write notes next to a git repo, and then ask where the SSH hosts are.

Two things already exist for the facts problem:

- [model](https://github.com/krisyotam/model): hosts, subagent CLIs,
  model slugs, a session log. Plain files under XDG paths.
- [prompt](https://github.com/krisyotam/prompt): the prompt / loop /
  graph artifacts a session was fed.

Neither of those says where *work* goes. That is machina.

## The failure mode

Left alone, agents:

- create `Downloads/`, `dev/`, `bin/`, `projects/`, `workspace/`
- put a clone next to a PDF in the same folder
- run `df -h` and misread the columns
- fill `/` because they wrote a dataset into `/tmp` or `$HOME`
- delete a directory they decided was "scratch"

The layout in TREE.md is the smallest set of names that covers how
work actually happens: versioned code, kept notes, experiments,
throwaways, an inbox. Then the media and study directories that
already exist on these machines.

## What "for agents" means

A human can keep a messy home and still find things. An agent reads
whatever `AGENTS.md` is in the directory it entered, and it believes
that file. So each work directory has one, short enough to read every
time, specific enough that the next action is obvious.

The check tools exist because the usual Unix commands are for humans
with a terminal. `disk` prints a table an agent can quote back without
guessing which column is available space.

## What stays out

Dotfiles stay in srice. Secrets stay in Infisical / `pass`. Per-project
rules stay in the project's own `AGENTS.md`. machina only claims the
home layout and the commands that inspect the box.
