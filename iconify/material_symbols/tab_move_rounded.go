package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabMoveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-2q0-.425.288-.713T4 16q.425 0 .713.288T5 17v2h14V7H5v2q0 .425-.288.713T4 10q-.425 0-.713-.288T3 9V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7.175-7H4q-.425 0-.713-.288T3 13q0-.425.288-.713T4 12h8.175L10.8 10.6q-.275-.275-.288-.688T10.8 9.2q.275-.275.7-.275t.7.275l3.1 3.1q.3.3.3.7t-.3.7l-3.1 3.1q-.275.275-.687.288T10.8 16.8q-.275-.275-.275-.7t.275-.7l1.375-1.4Z"/>`),
		g.Group(children),
	)
}