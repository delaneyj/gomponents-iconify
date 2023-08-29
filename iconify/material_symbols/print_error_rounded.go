package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintErrorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21q-.425 0-.713-.288T18 20q0-.425.288-.713T19 19q.425 0 .713.288T20 20q0 .425-.288.713T19 21Zm0-4q-.425 0-.713-.288T18 16v-3q0-.425.288-.713T19 12q.425 0 .713.288T20 13v3q0 .425-.288.713T19 17ZM7 21q-.425 0-.713-.288T6 20v-3H3q-.425 0-.713-.288T2 16v-5q0-1.275.875-2.138T5 8h14q1 0 1.763.563T21.825 10H18q-.825 0-1.413.588T16 12v3H8v4h8v2H7ZM6 7V5q0-.825.588-1.413T8 3h8q.825 0 1.413.588T18 5v2H6Z"/>`),
		g.Group(children),
	)
}