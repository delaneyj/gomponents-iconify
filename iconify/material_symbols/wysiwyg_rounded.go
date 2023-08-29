package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WysiwygRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V7H5v12Zm3-7q-.425 0-.713-.288T7 11q0-.425.288-.713T8 10h8q.425 0 .713.288T17 11q0 .425-.288.713T16 12H8Zm0 4q-.425 0-.713-.288T7 15q0-.425.288-.713T8 14h4q.425 0 .713.288T13 15q0 .425-.288.713T12 16H8Z"/>`),
		g.Group(children),
	)
}