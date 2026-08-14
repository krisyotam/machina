# Skills

Two skills ship in `skills/`. They are how a model that just landed
operates the box without reading this whole repo first.

| Skill | When |
|-------|------|
| `machina` | Session start, where a file goes, disk/mem/load |
| `machina-install` | Bootstrap or repair the layout |

The files in `skills/<name>/SKILL.md` are the source. Copies under
`.claude/skills` and `.grok/skills` in this checkout are symlinks
so a CLI opened on the repo finds them. `machina install skills`
writes the same files into each agent CLI's home skill directory:

```
~/.claude/skills/<name>/SKILL.md
~/.grok/skills/<name>/SKILL.md
~/.codex/skills/<name>/SKILL.md
~/.gemini/skills/<name>/SKILL.md
~/.config/opencode/skills/<name>/SKILL.md
~/.hermes/skills/<name>/SKILL.md
~/.agents/skills/<name>/SKILL.md
```

Those copies are ours. Install rewrites them. Do not edit a copy;
edit `skills/<name>/SKILL.md` and re-run `install skills`.

Directory purposes still live in `~/.config/machina/README.md` and
each `~/<dir>/AGENTS.md`. The skill points at those files. It does
not replace them.
