package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PropaneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 6h4V5h-4v1Zm6-1v1h1q2.5 0 4.25 1.75T23 12q0 2.5-1.75 4.25T17 18v2q0 .425-.288.713T16 21q-.425 0-.713-.288T15 20v-2H9v2q0 .425-.288.713T8 21q-.425 0-.713-.288T7 20v-2q-2.5 0-4.25-1.75T1 12q0-2.5 1.75-4.25T7 6h1V5q0-.825.588-1.413T10 3h4q.825 0 1.413.588T16 5Z"/>`),
		g.Group(children),
	)
}