package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13q-1.25 0-2.125-.875T11 10q0-1.25.875-2.125T14 7q1.25 0 2.125.875T17 10q0 1.25-.875 2.125T14 13Zm-7 3q-.825 0-1.413-.588T5 14V6q0-.825.588-1.413T7 4h14q.825 0 1.413.588T23 6v8q0 .825-.588 1.413T21 16H7Zm2-2h10q0-.825.588-1.413T21 12V8q-.825 0-1.413-.588T19 6H9q0 .825-.588 1.413T7 8v4q.825 0 1.413.588T9 14Zm-6 6q-.825 0-1.413-.588T1 18V8q0-.425.288-.713T2 7q.425 0 .713.288T3 8v10h16q.425 0 .713.288T20 19q0 .425-.288.713T19 20H3Zm4-6V6v8Z"/>`),
		g.Group(children),
	)
}