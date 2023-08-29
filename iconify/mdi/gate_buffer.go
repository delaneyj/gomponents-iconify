package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateBuffer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7.24L15.53 12L6 16.76V7.24M4 4v16l16-8L4 4Z"/>`),
		g.Group(children),
	)
}