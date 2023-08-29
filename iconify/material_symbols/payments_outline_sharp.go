package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13q-1.25 0-2.125-.875T11 10q0-1.25.875-2.125T14 7q1.25 0 2.125.875T17 10q0 1.25-.875 2.125T14 13Zm-9 3V4h18v12H5Zm4-2h10q0-.825.588-1.413T21 12V8q-.825 0-1.413-.588T19 6H9q0 .825-.588 1.413T7 8v4q.825 0 1.413.588T9 14Zm-8 6V7h2v11h17v2H1Zm6-6V6v8Z"/>`),
		g.Group(children),
	)
}