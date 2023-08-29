package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.942 22a.997.997 0 0 1-.707-.293l-4.628-4.628a1 1 0 0 1 0-1.414l4.628-4.63a1 1 0 0 1 1.414 1.415l-3.92 3.922l3.92 3.921A1 1 0 0 1 9.942 22Z"/><path fill="currentColor" d="M15.686 17.372H5.314a1 1 0 0 1 0-2h10.372a2.002 2.002 0 0 0 2-2V3a1 1 0 0 1 2 0v10.372a4.004 4.004 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}