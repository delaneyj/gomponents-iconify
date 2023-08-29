package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterPlot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21q-1.65 0-2.825-1.175T13 17q0-1.65 1.175-2.825T17 13q1.65 0 2.825 1.175T21 17q0 1.65-1.175 2.825T17 21ZM7 18q-1.65 0-2.825-1.175T3 14q0-1.65 1.175-2.825T7 10q1.65 0 2.825 1.175T11 14q0 1.65-1.175 2.825T7 18Zm4-8q-1.65 0-2.825-1.175T7 6q0-1.65 1.175-2.825T11 2q1.65 0 2.825 1.175T15 6q0 1.65-1.175 2.825T11 10Z"/>`),
		g.Group(children),
	)
}