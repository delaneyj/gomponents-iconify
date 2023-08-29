package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeBringForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h14v14H2V2m20 6v14H8v-4h2v2h10V10h-2V8h4Z"/>`),
		g.Group(children),
	)
}