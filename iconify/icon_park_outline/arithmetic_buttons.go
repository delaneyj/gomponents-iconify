package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArithmeticButtons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3ZM10 14h8m-8 0h8m-7 23l6-6m-3-13v-8m3 27l-6-6m19-17h8m-8 17h8m-8 6h8"/><path d="M24 4v40M4 24h40"/><path stroke-linejoin="round" d="M30 4H18m12 40H18M4 28v-8m40 8v-8"/></g>`),
		g.Group(children),
	)
}