package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 1H3a2 2 0 0 0-2 2v2h22V3a2 2 0 0 0-2-2ZM4 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3 0a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1Zm16 6V6H1v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}