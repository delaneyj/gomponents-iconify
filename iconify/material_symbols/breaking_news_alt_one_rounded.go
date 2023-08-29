package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakingNewsAltOneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h3q.425 0 .713-.288T11 16q0-.425-.288-.713T10 15H7q-.425 0-.713.288T6 16q0 .425.288.713T7 17Zm0-4h3q.425 0 .713-.288T11 12q0-.425-.288-.713T10 11H7q-.425 0-.713.288T6 12q0 .425.288.713T7 13Zm10-6q-.425 0-.713.288T16 8v4q0 .425.288.713T17 13q.425 0 .713-.288T18 12V8q0-.425-.288-.713T17 7ZM7 9h3q.425 0 .713-.288T11 8q0-.425-.288-.713T10 7H7q-.425 0-.713.288T6 8q0 .425.288.713T7 9ZM4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H4Zm13-4q.425 0 .713-.288T18 16q0-.425-.288-.713T17 15q-.425 0-.713.288T16 16q0 .425.288.713T17 17Z"/>`),
		g.Group(children),
	)
}