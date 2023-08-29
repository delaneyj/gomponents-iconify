package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PineTreeVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 12L12 2L5 12h1.86L3 18h7v4h4v-4h7l-3.86-6H19m-3.84-2H13.5l3.84 6H6.67l3.86-6H8.84L12 5.5l3.16 4.5Z"/>`),
		g.Group(children),
	)
}