# ~/dl

Downloads inbox. Files land here and leave here in the same
session.

Shared primitives (hosts/SSH, subagent CLIs, models.env) live in
`~/.config/model/`. Read `~/.config/model/README.md` on launch and
append a row to `~/.local/state/model/log.csv` when the session ends.

## Sort into

| Kind | Dest |
|------|------|
| A project tarball or clone | `~/src` or `~/lab` |
| A paper or dataset | `~/sci` |
| A book or text | `~/lit` |
| An image | `~/pics` |
| A video | `~/vids` |
| Audio | `~/music` |
| A note you will keep | `~/research` |
| Trash | delete |

## Rules

- Never create a directory named `Downloads` or `downloads`.
  This directory is the inbox.
- Never delete `~/dl`. Empty it; do not remove it.
- Do not build a folder structure in here. If a download is a tree,
  move the tree to its dest.
- Agents do not leave their own output in `~/dl`. Write output to
  `~/tmp/<date>-topic` or the dest directory.
