package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20v-1q-1.25 0-2.125-.875T1 16v-5q0-.825.588-1.413T3 9q.825 0 1.413.588T5 11v4h14v-4q0-.825.588-1.413T21 9q.825 0 1.413.588T23 11v5q0 1.25-.875 2.125T20 19v1q0 .425-.288.713T19 21q-.425 0-.713-.288T18 20v-1H6v1q0 .425-.288.713T5 21Zm2-8v-2q0-1.375-.838-2.463T4 7V6q0-1.25.875-2.125T7 3h10q1.25 0 2.125.875T20 6v1q-1.35.35-2.175 1.463T17 11v2H7Z"/>`),
		g.Group(children),
	)
}