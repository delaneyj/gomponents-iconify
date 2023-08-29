package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeTeal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="16" r="14" fill="#00aab4"/><path fill="#fff" d="M16.42 7.32h7v7h-7z"/>`),
		g.Group(children),
	)
}