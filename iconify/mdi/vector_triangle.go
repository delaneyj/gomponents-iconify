package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3v6h.73l-3.94 7H2v6h6v-2h8v2h6v-6h-3.79l-3.94-7H15V3m-4 2h2v2h-2m1 2.04l4 7.11V18H8v-1.85M4 18h2v2H4m14-2h2v2h-2"/>`),
		g.Group(children),
	)
}