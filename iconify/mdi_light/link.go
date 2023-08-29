package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 13v-1h7v1H8m7.5-6a5.5 5.5 0 0 1 5.5 5.5a5.5 5.5 0 0 1-5.5 5.5H13v-1h2.5c2.5 0 4.5-2 4.5-4.5S18 8 15.5 8H13V7h2.5m-8 11A5.5 5.5 0 0 1 2 12.5A5.5 5.5 0 0 1 7.5 7H10v1H7.5C5 8 3 10 3 12.5S5 17 7.5 17H10v1H7.5Z"/>`),
		g.Group(children),
	)
}