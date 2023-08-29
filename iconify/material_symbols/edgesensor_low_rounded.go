package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgesensorLowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 14q-.425 0-.713-.288T2 13V8q0-.425.288-.713T3 7q.425 0 .713.288T4 8v5q0 .425-.288.713T3 14Zm18 3q-.425 0-.713-.288T20 16v-5q0-.425.288-.713T21 10q.425 0 .713.288T22 11v5q0 .425-.288.713T21 17ZM8 22q-.825 0-1.413-.588T6 20V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v16q0 .825-.588 1.413T16 22H8Zm0-5h8V7H8v10Z"/>`),
		g.Group(children),
	)
}