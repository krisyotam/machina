# ~/lab

Experiments. One-off investigations. Things that might become a
repo or a kept note, and might not.

Shared primitives (hosts/SSH, subagent CLIs, models.env) live in
`~/.config/model/`. Read `~/.config/model/README.md` on launch and
append a row to `~/.local/state/model/log.csv` when the session ends.

## Put here

A named subdirectory for the experiment: `~/lab/<topic>/`. Give it
a real name, not `test` or `new`.

## Promote or delete

```
becomes a git repo     →  ~/src/<name>
becomes a kept write-up →  ~/research/<topic>
done and worthless      →  delete the subdirectory
```

Do not let lab become an unmarked junk drawer. If you cannot name
the subdirectory, it belongs in `~/tmp`.

## Put somewhere else

- A project you already know is a repo → `~/src`
- Notes you already know you will keep → `~/research`
- A file that should die with the session → `~/tmp`

## Rules

- Do not write files at the root of `~/lab`.
- Do not delete `~/lab` as a whole. Clean topic by topic, and only
  the topic you are sure about.
- Never create `bin/` or `dev/` in here.
- Do not build unless asked.
