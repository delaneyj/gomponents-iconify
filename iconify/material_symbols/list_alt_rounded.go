package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17q.425 0 .713-.288T9 16q0-.425-.288-.713T8 15q-.425 0-.713.288T7 16q0 .425.288.713T8 17Zm0-4q.425 0 .713-.288T9 12q0-.425-.288-.713T8 11q-.425 0-.713.288T7 12q0 .425.288.713T8 13Zm0-4q.425 0 .713-.288T9 8q0-.425-.288-.713T8 7q-.425 0-.713.288T7 8q0 .425.288.713T8 9Zm4 8h4q.425 0 .713-.288T17 16q0-.425-.288-.713T16 15h-4q-.425 0-.713.288T11 16q0 .425.288.713T12 17Zm0-4h4q.425 0 .713-.288T17 12q0-.425-.288-.713T16 11h-4q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm0-4h4q.425 0 .713-.288T17 8q0-.425-.288-.713T16 7h-4q-.425 0-.713.288T11 8q0 .425.288.713T12 9ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}