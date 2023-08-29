package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewTimelineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h4q.425 0 .713-.288T12 16q0-.425-.288-.713T11 15H7q-.425 0-.713.288T6 16q0 .425.288.713T7 17Zm3-4h4q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11h-4q-.425 0-.713.288T9 12q0 .425.288.713T10 13Zm3-4h4q.425 0 .713-.288T18 8q0-.425-.288-.713T17 7h-4q-.425 0-.713.288T12 8q0 .425.288.713T13 9ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}