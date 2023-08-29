package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilBarrelOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2v-6H3v-2h2V5H3V3h18v2h-2v6h2v2h-2v6h2v2H3Zm4-2h10v-6h-1v-2h1V5H7v6h1v2H7v6Zm5-3q1.25 0 2.125-.863T15 13.05q0-.975-.563-1.675T12 8.5q-1.875 2.15-2.438 2.863T9 13.05q0 1.225.875 2.087T12 16Zm-5 3V5v14Z"/>`),
		g.Group(children),
	)
}