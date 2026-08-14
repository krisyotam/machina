# Check tools

Agents should not scrape `df -h`, `free -h`, or `uptime`. Those
commands are for a human looking at a terminal. The numbers are in
the right columns if you already know the format, and in the wrong
ones if you are guessing.

`machina` prints small tables. The same binary answers to the short
names in `~/.local/bin`.

| Command | What it prints |
|---------|----------------|
| `disk` | real filesystems: device, size, used, avail, use%, mount |
| `mem` | total, used, available, swap |
| `load` | 1/5/15 minute load, cpu count |
| `home-map` | top-level home entries that machina knows, plus whether `AGENTS.md` is there |
| `machina where PATH` | which work dir a path belongs in |

`disk` warns on any filesystem at 90% full or above. It skips tmpfs,
proc, sysfs, cgroup, and other virtual mounts. Pass `-a` to see
those too.

`home-map` does not compute directory sizes by default. Walking
`~/src` can take minutes. Pass `-s` if you need sizes.

## Output rules

- ASCII. No color, no progress bars, no emoji.
- Stable column names so an agent can quote a line back.
- One filesystem or one fact per line.
- Exit `0` on success, `1` if the report itself failed, `2` on
  usage error. A full disk is still exit `0`; the warn line is the
  signal.

## Why these four

Disk is how agents notice they are about to write a dataset onto a
46G root. Memory is how they notice a runaway compile. Load is how
they notice the box is already busy. The home map is how they stop
asking where things go.

More tools can be added the same way: a function in `report.go`, a
name in `binLinks`. If the output needs a library to parse, the
command is the wrong shape.
