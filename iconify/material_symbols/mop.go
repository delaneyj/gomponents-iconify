package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.55 21H6v-2q0-.425.288-.713T7 18q.425 0 .713.288T8 19v2h3v-2q0-.425.288-.713T12 18q.425 0 .713.288T13 19v2h3v-2q0-.425.288-.713T17 18q.425 0 .713.288T18 19v2h2.45l-1-4H4.55l-1 4Zm16.9 2H3.55q-.975 0-1.575-.775t-.35-1.725L3 15v-2q0-.825.588-1.413T5 11h4V4q0-1.25.875-2.125T12 1q1.25 0 2.125.875T15 4v7h4q.825 0 1.413.588T21 13v2l1.375 5.5q.325.95-.288 1.725T20.45 23Z"/>`),
		g.Group(children),
	)
}