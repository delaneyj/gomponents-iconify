package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyYen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-4H6v-2h5v-2H6v-2h4.075L5 3h2.375L12 10.3L16.625 3H19l-5.075 8H18v2h-5v2h5v2h-5v4h-2Z"/>`),
		g.Group(children),
	)
}