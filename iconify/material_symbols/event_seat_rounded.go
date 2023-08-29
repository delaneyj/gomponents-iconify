package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventSeatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20v-3q0-.825.588-1.413T6 15h12q.825 0 1.413.588T20 17v3q0 .425-.288.713T19 21q-.425 0-.713-.288T18 20v-3H6v3q0 .425-.288.713T5 21Zm-.5-7q-.625 0-1.063-.438T3 12.5q0-.625.438-1.063T4.5 11q.625 0 1.063.438T6 12.5q0 .625-.438 1.063T4.5 14ZM7 14V5q0-.825.588-1.413T9 3h6q.825 0 1.413.588T17 5v9H7Zm12.5 0q-.625 0-1.063-.438T18 12.5q0-.625.438-1.063T19.5 11q.625 0 1.063.438T21 12.5q0 .625-.438 1.063T19.5 14Z"/>`),
		g.Group(children),
	)
}