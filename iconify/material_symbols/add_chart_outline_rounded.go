package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddChartOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h8q.425 0 .713.288T14 4q0 .425-.288.713T13 5H5v14h14v-8q0-.425.288-.713T20 10q.425 0 .713.288T21 11v8q0 .825-.588 1.413T19 21H5Zm3-11q-.425 0-.713.288T7 11v5q0 .425.288.713T8 17q.425 0 .713-.288T9 16v-5q0-.425-.288-.713T8 10Zm4-3q-.425 0-.713.288T11 8v8q0 .425.288.713T12 17q.425 0 .713-.288T13 16V8q0-.425-.288-.713T12 7Zm4 6q-.425 0-.713.288T15 14v2q0 .425.288.713T16 17q.425 0 .713-.288T17 16v-2q0-.425-.288-.713T16 13Zm1-6h-1q-.425 0-.713-.288T15 6q0-.425.288-.713T16 5h1V4q0-.425.288-.713T18 3q.425 0 .713.288T19 4v1h1q.425 0 .713.288T21 6q0 .425-.288.713T20 7h-1v1q0 .425-.288.713T18 9q-.425 0-.713-.288T17 8V7Zm-5 5Z"/>`),
		g.Group(children),
	)
}