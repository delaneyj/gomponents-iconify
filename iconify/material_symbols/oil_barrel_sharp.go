package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilBarrelSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2v-6H3v-2h2V5H3V3h18v2h-2v6h2v2h-2v6h2v2H3Zm9-5q1.25 0 2.125-.863T15 13.05q0-.975-.563-1.675T12 8.5q-1.875 2.15-2.438 2.863T9 13.05q0 1.225.875 2.087T12 16Z"/>`),
		g.Group(children),
	)
}