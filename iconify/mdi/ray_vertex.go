package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RayVertex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11h7.17c.41-1.17 1.52-2 2.83-2s2.42.83 2.83 2H22v2h-7.17A2.99 2.99 0 0 1 12 15a2.99 2.99 0 0 1-2.83-2H2v-2Z"/>`),
		g.Group(children),
	)
}