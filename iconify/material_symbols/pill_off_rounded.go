package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.35 14.6L9.4 6.65l2-2q.8-.8 1.825-1.225T15.375 3q2.35 0 3.988 1.638T21 8.625q0 1.125-.425 2.15T19.35 12.6l-2 2Zm-2.825 2.825L12.6 19.35q-.8.8-1.825 1.225T8.625 21q-2.35 0-3.988-1.637T3 15.375q0-1.125.425-2.15T4.65 11.4l1.925-1.925L2.1 5q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275T19.1 22l-4.575-4.575Z"/>`),
		g.Group(children),
	)
}