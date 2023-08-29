package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatPumpBalanceOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 18.6H4v-12h2v10h2v-10h6v10h2v-10h4.175L19 5.425L20.4 4L24 7.6l-3.6 3.575L19 9.75l1.175-1.15H18v10h-6v-10h-2v10Zm-9 3v-10h22v10H1Zm2-2h18v-6H3v6Zm18-6H3h18Z"/>`),
		g.Group(children),
	)
}