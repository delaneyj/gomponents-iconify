package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18v2h-6.5A6.5 6.5 0 0 1 7 13.5V7.83l-3.09 3.09L2.5 9.5L8 4l5.5 5.5l-1.41 1.41L9 7.83v5.67C9 16 11 18 13.5 18H20Z"/>`),
		g.Group(children),
	)
}