# ~/tmp

Disposable. Wipe-safe. If you would mourn the file, it does not
belong here.

Shared primitives (hosts/SSH, subagent CLIs, models.env) live in
`~/.config/model/`. Read `~/.config/model/README.md` on launch and
append a row to `~/.local/state/model/log.csv` when the session ends.

## Naming

```
~/tmp/<YYYY-MM-DD>-<topic>/
```

Example: `~/tmp/2026-08-14-pdf-extract/`. Work inside that
directory, not at the root of `~/tmp`.

## Rules

- No secrets. No keys, no `models.env`, no mail.
- No clones of real projects. Use `~/lab` or `~/src`.
- Delete the dated directory when the session no longer needs it.
- A dated directory still sitting here a week later should be
  promoted (`~/lab`, `~/research`, `~/src`) or removed.
- This is the only home directory machina will call wipe-safe.
  `rm -r` of a dated child is fine. `rm -r ~/tmp` itself is a
  human decision.

## Put somewhere else

- Kept notes → `~/research`
- Experiments you might return to → `~/lab`
- A real project → `~/src`
- A download you have not sorted → `~/dl`
