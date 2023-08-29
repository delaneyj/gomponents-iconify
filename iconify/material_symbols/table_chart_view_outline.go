package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableChartViewOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.4 22L6 20.6l6.9-6.925l3.5 3.5L21.575 12L23 13.425L16.4 20l-3.5-3.5L7.4 22ZM4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h14q.825 0 1.413.588T20 5v4.2H4V21ZM4 7.2h14V5H4v2.2Zm0 0V5v2.2Z"/>`),
		g.Group(children),
	)
}