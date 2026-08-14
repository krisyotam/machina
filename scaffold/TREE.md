# Tree

This is the spec. Every path `machina install` writes or touches
lives here. `machina tree` prints this file.

Paths are relative to `$HOME`. Existing directories stay. Existing
files are not overwritten unless noted. `hosts.md` and `models.env`
are never overwritten.

```
~/
  .config/
    model/                            from github.com/krisyotam/model
      README.md                       index of shared facts
      hosts.md                        machines, SSH, sudo   (never overwrite)
      subagents.md                    how to invoke other agent CLIs
      models.md                       slug index
      models.env                      shared env, chmod 600 (never overwrite)
    machina/
      README.md                       machine index (rewritten on install)
    opencode/
      AGENTS.md                       pointer line appended
      skills/
        machina/SKILL.md              operate the box
        machina-install/SKILL.md      bootstrap or repair

  .local/
    state/
      model/
        log.csv                       session log, one row per session
    bin/
      machina                         this program
      disk -> machina                 filesystems
      mem  -> machina                 memory
      load -> machina                 load averages
      home-map -> machina             top-level home map

  .claude/
    CLAUDE.md                         pointer line appended
    skills/
      machina/SKILL.md
      machina-install/SKILL.md
  .codex/
    AGENTS.md                         pointer line appended
    skills/
      machina/SKILL.md
      machina-install/SKILL.md
  .gemini/
    GEMINI.md                         pointer line appended
    skills/
      machina/SKILL.md
      machina-install/SKILL.md
  .grok/
    AGENTS.md                         pointer line appended
    skills/
      machina/SKILL.md
      machina-install/SKILL.md
  .hermes/
    SOUL.md                           pointer line appended
    skills/
      machina/SKILL.md
      machina-install/SKILL.md
  .agents/
    skills/
      machina/SKILL.md
      machina-install/SKILL.md

  src/
    AGENTS.md                         git repos, one project per directory
  research/
    AGENTS.md                         kept notes and investigations
  lab/
    AGENTS.md                         experiments that are not repos yet
  tmp/
    AGENTS.md                         disposable work, dated subdirs
  dl/
    AGENTS.md                         downloads inbox
  edu/
    AGENTS.md                         coursework and lectures
  lit/
    AGENTS.md                         books and texts
  sci/
    AGENTS.md                         papers and datasets
  pics/
    AGENTS.md                         images
  vids/
    AGENTS.md                         video
  music/
    AGENTS.md                         audio
```

## What each band is

| Band | Path | Who owns it |
|------|------|-------------|
| Facts | `.config/model/` | [krisyotam/model](https://github.com/krisyotam/model). machina copies the templates. A human fills `hosts.md` and `models.env`. |
| Index | `.config/machina/README.md` | machina. Rewritten on install. |
| State | `.local/state/model/log.csv` | every agent, one row per session |
| Tools | `.local/bin/` | `machina` and the four links |
| Glue | `.claude/`, `.codex/`, `.gemini/`, `.grok/`, `.hermes/`, `.config/opencode/`, `.agents/` | pointer line in the instruction file; the two skills copied in |
| Work | `src` `research` `lab` `tmp` `dl` `edu` `lit` `sci` `pics` `vids` `music` | `AGENTS.md` written if missing (`-f` to replace) |

Skill files are ours. Install rewrites them. Edit
`skills/<name>/SKILL.md` in this repo, then `machina install skills`.

## Not in this tree

These exist next to machina. Install does not create them.

```
~/.prompt/          prompt / loop / graph artifacts (krisyotam/prompt)
~/.cache/machina/   clone of krisyotam/model if ~/src/model is missing
```

These names must not appear:

```
~/bin               scripts go in ~/.local/bin
~/dev               code goes in ~/src
~/Downloads         inbox is ~/dl
```

## Rules the tree encodes

- Work goes in a named directory. `$HOME` is not a scratch pad.
- `src` is git repos. Loose files do not live there.
- `research` is kept writing. `lab` is an experiment. `tmp` is trash
  with a date.
- `dl` is an inbox. Sort out of it in the same session.
- Protected directories are never deleted: `src`, `dl`, `edu`, `lit`,
  `music`, `pics`, `sci`, `vids`.
- `tmp` is the only wipe-safe directory.

[docs/layout.md](docs/layout.md) is what belongs in each work
directory. [docs/guards.md](docs/guards.md) is the hard blocks.
[docs/install.md](docs/install.md) is overwrite and dry-run.
