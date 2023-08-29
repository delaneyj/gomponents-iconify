package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-9q0-1.875.713-3.513t1.924-2.85q1.213-1.212 2.85-1.924T12 3q1.875 0 3.513.713t2.85 1.924q1.212 1.213 1.925 2.85T21 12v9h-6v-8h4v-1q0-2.925-2.038-4.963T12 5Q9.075 5 7.037 7.038T5 12v1h4v8H3Z"/>`),
		g.Group(children),
	)
}