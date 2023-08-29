package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PropaneTank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 13v-3q0-1.425.85-2.475T7 6.125V4q0-.825.588-1.413T9 2h6q.825 0 1.413.588T17 4v2.125q1.3.35 2.15 1.4T20 10v3H4Zm4 9q-1.65 0-2.825-1.175T4 18v-3h16v3q0 1.65-1.175 2.825T16 22H8ZM9 6h2q0-.425.288-.713T12 5q.425 0 .713.288T13 6h2V4H9v2Z"/>`),
		g.Group(children),
	)
}