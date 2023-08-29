package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weekend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-1.25 0-2.125-.875T1 17v-5q0-.825.588-1.413T3 10q.825 0 1.413.588T5 12v4h14v-4q0-.825.588-1.413T21 10q.825 0 1.413.588T23 12v5q0 1.25-.875 2.125T20 20H4Zm3-6v-2q0-1.325-.863-2.325T4 8.3V7q0-1.25.875-2.125T7 4h10q1.25 0 2.125.875T20 7v1.3q-1.325.275-2.163 1.313T17 12v2H7Z"/>`),
		g.Group(children),
	)
}