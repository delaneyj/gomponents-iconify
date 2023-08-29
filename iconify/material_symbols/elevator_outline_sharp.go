package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18h3v-4h1V9.5H6V14h1v4Zm1.5-9.5q.525 0 .888-.363t.362-.887q0-.525-.363-.888T8.5 6q-.525 0-.888.363t-.362.887q0 .525.363.888T8.5 8.5ZM13 11h5l-2.5-4l-2.5 4Zm2.5 6l2.5-4h-5l2.5 4ZM3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}