package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18v-3q0-1.65 1.175-2.825T6 11h9v-1q0-.425-.288-.713T14 9h-4q-.425 0-.713.288T9 10H7q0-1.25.875-2.125T10 7h4q1.25 0 2.125.875T17 10v4q.425 0 .713-.288T18 13V9q0-1.25.875-2.125T21 6h1v2h-1q-.425 0-.713.288T20 9v4q0 1.25-.875 2.125T17 16v2H2Zm2-2h11v-3H6q-.825 0-1.413.588T4 15v1Zm11 0v-3v3Z"/>`),
		g.Group(children),
	)
}