package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventSeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-6h16v6h-2v-4H6v4H4Zm.5-7q-.625 0-1.063-.438T3 12.5q0-.625.438-1.063T4.5 11q.625 0 1.063.438T6 12.5q0 .625-.438 1.063T4.5 14ZM7 14V5q0-.825.588-1.413T9 3h6q.825 0 1.413.588T17 5v9H7Zm12.5 0q-.625 0-1.063-.438T18 12.5q0-.625.438-1.063T19.5 11q.625 0 1.063.438T21 12.5q0 .625-.438 1.063T19.5 14Z"/>`),
		g.Group(children),
	)
}