# ~/pics

Images. The library, not a staging area.

Shared primitives (hosts/SSH, subagent CLIs, models.env) live in
`~/.config/model/`. Read `~/.config/model/README.md` on launch and
append a row to `~/.local/state/model/log.csv` when the session ends.

## Rules

- Never delete `~/pics`.
- Do not reorganize unless asked.
- New images you generated or downloaded this session go to
  `~/tmp/<YYYY-MM-DD>-topic` or `~/dl` first. Move them here only
  when they are keepers.
- Site images for krisyotam.com do not live here. They go to
  stargate CDN storage. See `~/.claude/docs/cdn.md` when that
  applies.
