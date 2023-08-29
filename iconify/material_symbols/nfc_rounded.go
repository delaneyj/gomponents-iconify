package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NfcRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 17q.425 0 .713-.288T17 16V8q0-.425-.288-.713T16 7h-3q-.825 0-1.413.588T11 9v1.3q-.5.275-.75.7T10 12q0 .825.588 1.413T12 14q.825 0 1.413-.588T14 12q0-.575-.275-1T13 10.3V9h2v6H9V9q.425 0 .713-.288T10 8q0-.425-.288-.713T9 7H8q-.425 0-.713.288T7 8v8q0 .425.288.713T8 17h8ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}