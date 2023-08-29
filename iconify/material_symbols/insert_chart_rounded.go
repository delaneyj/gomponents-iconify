package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17q.425 0 .713-.288T9 16v-5q0-.425-.288-.713T8 10q-.425 0-.713.288T7 11v5q0 .425.288.713T8 17Zm4 0q.425 0 .713-.288T13 16V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v8q0 .425.288.713T12 17Zm4 0q.425 0 .713-.288T17 16v-2q0-.425-.288-.713T16 13q-.425 0-.713.288T15 14v2q0 .425.288.713T16 17ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}