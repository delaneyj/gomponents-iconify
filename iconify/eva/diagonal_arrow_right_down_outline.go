package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagonalArrowRightDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaDiagonalArrowRightDownOutline0"><g id="evaDiagonalArrowRightDownOutline1"><path id="evaDiagonalArrowRightDownOutline2" fill="currentColor" d="M17 8a1 1 0 0 0-1 1v5.59l-8.29-8.3a1 1 0 0 0-1.42 1.42l8.3 8.29H9a1 1 0 0 0 0 2h8a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1Z"/></g></g>`),
		g.Group(children),
	)
}