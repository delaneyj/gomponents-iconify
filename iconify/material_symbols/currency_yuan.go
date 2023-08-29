package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyYuan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-7H6v-2h4.725L5 3h2.375L12 10.3L16.625 3H19l-5.725 9H18v2h-5v7h-2Z"/>`),
		g.Group(children),
	)
}