package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateAnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4v16h12a8 8 0 0 0 8-8a8 8 0 0 0-8-8H2m2 2h10a6 6 0 0 1 6 6a6 6 0 0 1-6 6H4V6Z"/>`),
		g.Group(children),
	)
}