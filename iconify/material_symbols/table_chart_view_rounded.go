package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableChartViewRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.7 21.3q-.275-.275-.275-.7t.275-.7l4.775-4.8q.575-.575 1.425-.575t1.425.575l2.075 2.075l4.475-4.475q.3-.3.7-.3t.7.3q.3.3.3.713t-.3.712l-4.45 4.475q-.575.575-1.413.575T15 18.6l-2.1-2.1l-4.8 4.8q-.275.275-.688.288T6.7 21.3ZM4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h14q.825 0 1.413.588T20 5v4q0 .425-.288.713T19 10H4v11Z"/>`),
		g.Group(children),
	)
}