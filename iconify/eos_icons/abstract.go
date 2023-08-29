package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abstract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 3H8a2 2 0 0 0-1.71 1l-4 7a2 2 0 0 0 0 2l4 7A2 2 0 0 0 8 21h8a2 2 0 0 0 1.74-1l4-7a2 2 0 0 0 0-2l-4-7A2 2 0 0 0 16 3Z"/>`),
		g.Group(children),
	)
}