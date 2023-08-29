package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13q1.25 0 2.125-.875T17 10q0-1.25-.875-2.125T14 7q-1.25 0-2.125.875T11 10q0 1.25.875 2.125T14 13Zm-9 3V4h18v12H5Zm-4 4V7h2v11h17v2H1ZM7 8q.825 0 1.413-.588T9 6H7v2Zm14 0V6h-2q0 .825.588 1.413T21 8Zm-2 6h2v-2q-.825 0-1.413.588T19 14ZM7 14h2q0-.825-.588-1.413T7 12v2Z"/>`),
		g.Group(children),
	)
}