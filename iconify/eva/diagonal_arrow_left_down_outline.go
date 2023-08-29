package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagonalArrowLeftDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaDiagonalArrowLeftDownOutline0"><g id="evaDiagonalArrowLeftDownOutline1"><path id="evaDiagonalArrowLeftDownOutline2" fill="currentColor" d="M17.71 6.29a1 1 0 0 0-1.42 0L8 14.59V9a1 1 0 0 0-2 0v8a1 1 0 0 0 1 1h8a1 1 0 0 0 0-2H9.41l8.3-8.29a1 1 0 0 0 0-1.42Z"/></g></g>`),
		g.Group(children),
	)
}