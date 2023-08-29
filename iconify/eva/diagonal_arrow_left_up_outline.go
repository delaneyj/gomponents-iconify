package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagonalArrowLeftUpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaDiagonalArrowLeftUpOutline0"><g id="evaDiagonalArrowLeftUpOutline1"><path id="evaDiagonalArrowLeftUpOutline2" fill="currentColor" d="M17.71 16.29L9.42 8H15a1 1 0 0 0 0-2H7.05a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1H7a1 1 0 0 0 1-1V9.45l8.26 8.26a1 1 0 0 0 1.42 0a1 1 0 0 0 .03-1.42Z"/></g></g>`),
		g.Group(children),
	)
}