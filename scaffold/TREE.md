# Tree

The layout `machina install` writes. Paths are relative to `$HOME`.
Directories already present are left in place. Files already present
are not overwritten unless `-f` is set (and even then, `hosts.md` and
`models.env` are never overwritten).

```
~/
  .config/
    model/                      from github.com/krisyotam/model
      README.md                 index of shared facts
      hosts.md                  machines, SSH, sudo
      subagents.md              how to invoke other agent CLIs
      models.md                 slug index
      models.env                shared env, chmod 600, not in git
    machina/
      README.md                 machine index (this layout)
  .local/
    state/
      model/
        log.csv                 session log, one row per session
    bin/
      machina                   this program
      disk -> machina           filesystems
      mem  -> machina           memory
      load -> machina           load averages
      home-map -> machina       top-level home map
  src/
    AGENTS.md                   versioned projects, one repo each
  research/
    AGENTS.md                   kept notes and investigations
  lab/
    AGENTS.md                   experiments that are not repos yet
  tmp/
    AGENTS.md                   disposable work, dated subdirs
  dl/
    AGENTS.md                   downloads inbox
  edu/
    AGENTS.md                   coursework and lectures
  lit/
    AGENTS.md                   books and texts
  sci/
    AGENTS.md                   papers and datasets
  pics/
    AGENTS.md                   images
  vids/
    AGENTS.md                   video
  music/
    AGENTS.md                   audio
```

Sibling standards, not installed by machina:

```
  .config/model/     facts every agent reads on launch
  .prompt/           prompt / loop / graph artifacts (krisyotam/prompt)
```

## Rules the tree encodes

- Work goes in a named directory. The root of `$HOME` is not a scratch pad.
- `src` is git repos. Loose files do not live there.
- `research` is kept writing. `lab` is an experiment. `tmp` is trash with a date.
- `dl` is an inbox. Sort out of it in the same session.
- Scripts live in `~/.local/bin`. There is no `~/bin` and no `bin/` directory
  anywhere this tree creates.
- There is no `~/dev` and no `~/Downloads`. Those names are banned.
- Protected directories are never deleted: `src`, `dl`, `edu`, `lit`,
  `music`, `pics`, `sci`, `vids`.

See [docs/layout.md](docs/layout.md) for what belongs in each directory
and [docs/guards.md](docs/guards.md) for the hard blocks.
