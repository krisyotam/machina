// Package skills holds the agent skills machina installs into each
// CLI's skill directory. The SKILL.md files in this package are the
// source of truth; copies under .claude/skills and .grok/skills in
// this repo are symlinks.
package skills

import "embed"

// FS is the skill tree. Paths: <name>/SKILL.md
//
//go:embed machina/SKILL.md
//go:embed machina-install/SKILL.md
var FS embed.FS

// Names is the install set. Keep in sync with the embed directives.
var Names = []string{"machina", "machina-install"}
