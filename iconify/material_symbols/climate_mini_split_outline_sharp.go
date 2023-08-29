package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClimateMiniSplitOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12V2h20v10H2Zm2 7v-2q1.25 0 2.125-.875T7 14h2q0 2.075-1.463 3.538T4 19Zm16 0q-2.075 0-3.538-1.463T15 14h2q0 1.25.875 2.125T20 17v2Zm-9 1v-6h2v6h-2Zm7-10h2H4h14ZM6 10V6h12v4h-2V8H8v2H6Zm-2 0h16V4H4v6Z"/>`),
		g.Group(children),
	)
}