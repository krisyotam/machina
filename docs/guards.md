# Guards

These are hard blocks. An agent that hits one stops and says so.
They are also written into the directory `AGENTS.md` files so a
session that never read this page still sees them.

## Do not delete

These directories are permanently protected. `rm`, `rm -r`,
`rm -rf`, moving them to trash, or deleting them as a side effect
of deleting a parent is forbidden.

```
~/dl   ~/edu   ~/lit   ~/music   ~/pics   ~/sci   ~/src   ~/vids
```

If a command would delete one of these, the reply is:

```
unauthorized access: LLMs are expressly forbidden from attempting to delete this route.
```

## Do not create

- A new top-level file or directory in `$HOME` without `sudo`.
  `machina install` is the exception, and only for the names in
  TREE.md.
- A directory named `bin` anywhere. Scripts go in `~/.local/bin`.
- A directory named `dev` anywhere under `$HOME`. Code lives in
  `~/src`.
- A directory named `Downloads` or `downloads` anywhere. The inbox
  is `~/dl`.

## Do not run

- `remaps` or `remapd`. Caps Lock → Super is an XKB option
  (`caps:super`). remapd has pegged Xorg at 70% CPU.
- `make`, `sudo make install`, `npm run build`, `cargo build`, or
  any compile unless the human said "build" in the message that
  asked for the work.
- `kill`, `pkill`, `killall`, `kill -9` unless the human named that
  process in the same conversation.

## Do not touch

- `~/.config/model/models.env` values go in via a secret manager,
  not by an agent inventing keys.
- `~/.config/model/hosts.md` is facts. Correct a wrong address if
  you verified it. Do not add a second copy of those facts to any
  other file.
- Production on stargate. See `hosts.md`. Restarting Docker or
  rewriting nginx there is a human decision.

## Why they live here

Guards that only exist in a chat system prompt get forgotten the
moment the agent `cd`s into a directory and reads a local file.
Putting them in the directory files is how they survive a fresh
context.
