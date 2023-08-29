package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PropaneTankOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-1.65 0-2.825-1.175T4 18v-8q0-1.425.85-2.475T7 6.125V4q0-.825.588-1.413T9 2h6q.825 0 1.413.588T17 4v2.125q1.3.35 2.15 1.4T20 10v8q0 1.65-1.175 2.825T16 22H8Zm-2-9h12v-3q0-.825-.588-1.413T16 8H8q-.825 0-1.413.588T6 10v3Zm2 7h8q.825 0 1.413-.588T18 18v-3H6v3q0 .825.588 1.413T8 20Zm4-7Zm0 2Zm0-1Zm1-8h2V4H9v2h2q0-.425.288-.713T12 5q.425 0 .713.288T13 6Z"/>`),
		g.Group(children),
	)
}