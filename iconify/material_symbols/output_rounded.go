package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutputRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v2h-2V5H5v14h14v-2h2v2q0 .825-.588 1.413T19 21H5Zm13.175-8H10q-.425 0-.713-.288T9 12q0-.425.288-.713T10 11h8.175L16.3 9.1q-.275-.275-.288-.688T16.3 7.7q.275-.275.7-.275t.7.275l3.6 3.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-3.6 3.6q-.275.275-.688.288T16.3 16.3q-.275-.275-.275-.7t.275-.7l1.875-1.9Z"/>`),
		g.Group(children),
	)
}