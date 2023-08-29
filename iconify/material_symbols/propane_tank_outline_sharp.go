package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PropaneTankOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-1.65 0-2.825-1.175T4 18v-8q0-1.425.85-2.475T7 6.125V2h10v4.125q1.3.35 2.15 1.4T20 10v8q0 1.65-1.175 2.825T16 22H8ZM9 6h6V4H9v2Zm2 0q0-.425.288-.713T12 5q.425 0 .713.288T13 6h-2Zm-5 7h12v-3q0-.825-.588-1.413T16 8H8q-.825 0-1.413.588T6 10v3Zm2 7h8q.825 0 1.413-.588T18 18v-3H6v3q0 .825.588 1.413T8 20Zm4-7Zm0 2Zm0-1Z"/>`),
		g.Group(children),
	)
}