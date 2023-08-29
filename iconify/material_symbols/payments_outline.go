package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13q-1.25 0-2.125-.875T11 10q0-1.25.875-2.125T14 7q1.25 0 2.125.875T17 10q0 1.25-.875 2.125T14 13Zm-7 3q-.825 0-1.413-.588T5 14V6q0-.825.588-1.413T7 4h14q.825 0 1.413.588T23 6v8q0 .825-.588 1.413T21 16H7Zm2-2h10q0-.825.588-1.413T21 12V8q-.825 0-1.413-.588T19 6H9q0 .825-.588 1.413T7 8v4q.825 0 1.413.588T9 14Zm11 6H3q-.825 0-1.413-.588T1 18V7h2v11h17v2ZM7 14V6v8Z"/>`),
		g.Group(children),
	)
}