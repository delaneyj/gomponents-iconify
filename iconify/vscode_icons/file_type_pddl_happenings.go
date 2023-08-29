package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypePddlHappenings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#0055d4" d="M11.36 4.15h17.37v6.45H11.36z"/><ellipse cx="5.863" cy="7.375" fill="#0055d4" rx="2.638" ry="2.785"/><path fill="#0055d4" d="M11.433 12.945h17.37v6.45h-17.37z"/><ellipse cx="5.936" cy="16.169" fill="#0055d4" rx="2.638" ry="2.785"/><path fill="#0055d4" d="M11.433 21.739h17.37v6.45h-17.37z"/><ellipse cx="5.936" cy="24.964" fill="#0055d4" rx="2.638" ry="2.785"/>`),
		g.Group(children),
	)
}