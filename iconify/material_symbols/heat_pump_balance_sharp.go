package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatPumpBalanceSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18.6h6v-10h2v10h6v-10h2.175L19 9.75l1.4 1.425L24 7.6L20.4 4L19 5.425L20.175 6.6H16v10h-2v-10H8v10H6v-10H4v12Zm-3 3v-10h22v10H1Z"/>`),
		g.Group(children),
	)
}