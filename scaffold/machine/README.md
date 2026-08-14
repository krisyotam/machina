# machina

Machine layout for agents on this box. Shared facts (hosts, subagent
CLIs, models.env) live in `~/.config/model/`. Read
`~/.config/model/README.md` on launch and append a row to
`~/.local/state/model/log.csv` when the session ends.

This file is the index of *where work goes*. Each work directory has
its own `AGENTS.md`. Closer file wins.

## Directories

| Path | Purpose |
|------|---------|
| `~/src` | Git repos. One project per directory. |
| `~/research` | Kept notes and investigations. |
| `~/lab` | Experiments that are not repos yet. |
| `~/tmp` | Disposable work. Dated subdirs. Wipe-safe. |
| `~/dl` | Downloads inbox. Sort out the same session. |
| `~/edu` | Coursework and lectures. |
| `~/lit` | Books and texts. |
| `~/sci` | Papers and datasets. |
| `~/pics` | Images. Do not reorganize unless asked. |
| `~/vids` | Video. Do not reorganize unless asked. |
| `~/music` | Audio. Do not reorganize unless asked. |

`machina where PATH` classifies a path into one of these.

## Check tools

Run these from `~/.local/bin` instead of scraping `df` / `free` /
`uptime`:

```
disk        filesystems, sizes, percent full
mem         memory and swap
load        load averages, cpu count
home-map    the table above, plus whether AGENTS.md is present
```

## Guards

- Never delete `~/src`, `~/dl`, `~/edu`, `~/lit`, `~/music`,
  `~/pics`, `~/sci`, `~/vids`.
- Never create `bin/` anywhere, `~/dev`, or `Downloads/`.
- Never run `remaps` or `remapd`.
- Never build (`make`, `cargo build`, `npm run build`) unless asked.
- Never kill a process unless it was named in the same conversation.
- Do not create a new top-level name under `$HOME`. The names above
  are the set.

## Skills

`machina` (operate the box) and `machina-install` (bootstrap or
repair) live in each CLI's skill directory after install. The
source is `~/src/machina/skills/`.

## Sibling standards

- `~/.config/model/`: facts every agent reads on launch
- `~/.prompt/`: prompt / loop / graph artifacts
- this file: home layout
