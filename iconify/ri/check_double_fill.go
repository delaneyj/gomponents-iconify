package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckDoubleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.602 13.76l1.412 1.412l8.466-8.466l1.414 1.415l-9.88 9.88l-6.364-6.365l1.414-1.414l2.125 2.125l1.413 1.412Zm.002-2.828l4.952-4.953l1.41 1.41l-4.952 4.953l-1.41-1.41Zm-2.827 5.655L7.364 18L1 11.636l1.414-1.414l1.413 1.413l-.001.001l4.951 4.951Z"/>`),
		g.Group(children),
	)
}