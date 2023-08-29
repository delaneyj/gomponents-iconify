package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeBringToFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7V5H4v5h2v1H3V4h7v3H9m4 14v-3h1v2h5v-5h-2v-1h3v7h-7M8 9h7v7H8V9m1 1v5h5v-5H9Z"/>`),
		g.Group(children),
	)
}