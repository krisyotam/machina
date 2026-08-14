# Layout

Every work directory is a name and a "do not". Agents read
`~/<dir>/AGENTS.md` before writing into that directory.

| Dir | Purpose | Do not |
|-----|---------|--------|
| `src` | Git repos. One project per directory. | Loose files, notes, downloads |
| `research` | Kept notes, write-ups, literature reviews | Clones, throwaway scripts, PDFs dumped at the root |
| `lab` | Experiments that are not repos yet | Finished write-ups (promote to research), real projects (promote to src) |
| `tmp` | Disposable work in dated subdirs | Anything you would mourn; secrets |
| `dl` | Downloads inbox | Long-term storage; a directory named `Downloads` |
| `edu` | Coursework, lectures, assigned reading | Personal research notes |
| `lit` | Books and texts | Notes *about* books (those go in research) |
| `sci` | Papers and datasets | Write-ups of papers (those go in research) |
| `pics` | Images | Reorganizing unless asked |
| `vids` | Video | Reorganizing unless asked |
| `music` | Audio | Reorganizing unless asked |

## How work moves

```
dl  --sort-->  the dest directory
lab --if it becomes a repo-->  src/<name>
lab --if it becomes a kept note-->  research/<topic>
tmp --when the session ends-->  delete
```

`tmp` uses dated names: `~/tmp/2026-08-14-topic`. If a dated dir is
still there a week later, it should have been promoted or deleted.

## Protected directories

These exist to be kept. An agent never deletes them, never `mv`s them
to trash, and never suggests wiping them to free space:

`src`, `dl`, `edu`, `lit`, `music`, `pics`, `sci`, `vids`

`tmp` is the only directory machina will describe as wipe-safe.
`lab` can be cleaned by hand, topic by topic, not as a unit.

## Names that must not appear

- `bin/` anywhere. Scripts go in `~/.local/bin`.
- `dev/` under `$HOME`. Development is `~/src`.
- `Downloads/` or `downloads/`. The inbox is `~/dl`.

Creating a new top-level name under `$HOME` is a decision, not a
convenience. On Kris's machines it needs `sudo` if you are an agent.
`machina install` is the one program allowed to create the names in
TREE.md.

## XDG

Config and state follow the [XDG Base Directory][xdg] spec.

| Kind | Path | Holds |
|------|------|-------|
| Config | `$XDG_CONFIG_HOME/model/` | Shared facts (model standard) |
| Config | `$XDG_CONFIG_HOME/machina/` | Machine index |
| State | `$XDG_STATE_HOME/model/` | `log.csv` |
| Exec | `~/.local/bin/` | `machina` and the check-tool links |

If `XDG_CONFIG_HOME` or `XDG_STATE_HOME` is unset, the defaults are
`~/.config` and `~/.local/state`.

[xdg]: https://specifications.freedesktop.org/basedir-spec/latest/
