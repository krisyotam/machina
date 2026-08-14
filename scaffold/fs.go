// Package scaffold holds the templates machina writes onto $HOME.
package scaffold

import "embed"

// FS is the template tree. Paths:
//
//	TREE.md
//	machine/README.md   →  ~/.config/machina/README.md
//	dirs/<name>/AGENTS.md  →  ~/<name>/AGENTS.md
//
//go:embed TREE.md
//go:embed machine/README.md
//go:embed all:dirs
var FS embed.FS
