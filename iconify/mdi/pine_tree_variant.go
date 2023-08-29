package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PineTreeVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 12L12 2L5 12h1.86L3 18h7v4h4v-4h7l-3.86-6H19Z"/>`),
		g.Group(children),
	)
}