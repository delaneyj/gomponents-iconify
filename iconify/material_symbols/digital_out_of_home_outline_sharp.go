package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalOutOfHomeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 23V1h11v22H9v-2H2v2H0Zm15.5 0v-6.6l2.15-2.05l-.525-2.6q-.975 1.125-2.363 1.688T12 14v-2q1.2 0 2.325-.575t1.85-1.75l.75-1.225q.375-.625 1.1-.85t1.375.05L24 9.6v4.9h-2v-3.575l-1.425-.6L23 23h-2.05l-1.525-7.15l-1.925 1.8V23h-2ZM2 19h7V3H2v16Zm2-6l3.5-2L4 9v4Zm13-6q-.825 0-1.413-.588T15 5q0-.825.588-1.413T17 3q.825 0 1.413.588T19 5q0 .825-.588 1.413T17 7ZM2 19h7h-7Z"/>`),
		g.Group(children),
	)
}