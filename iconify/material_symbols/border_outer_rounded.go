package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOuterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h14V5H5v14Zm-2 0V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5q-.825 0-1.413-.588T3 19Zm9-10q-.425 0-.713-.288T11 8q0-.425.288-.713T12 7q.425 0 .713.288T13 8q0 .425-.288.713T12 9Zm-4 4q-.425 0-.713-.288T7 12q0-.425.288-.713T8 11q.425 0 .713.288T9 12q0 .425-.288.713T8 13Zm4 0q-.425 0-.713-.288T11 12q0-.425.288-.713T12 11q.425 0 .713.288T13 12q0 .425-.288.713T12 13Zm4 0q-.425 0-.713-.288T15 12q0-.425.288-.713T16 11q.425 0 .713.288T17 12q0 .425-.288.713T16 13Zm-4 4q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17Z"/>`),
		g.Group(children),
	)
}