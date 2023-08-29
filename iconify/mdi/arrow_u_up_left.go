package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6v7.5a6.5 6.5 0 1 1-13 0V7.83l-3.09 3.09L2.5 9.5L8 4l5.5 5.5l-1.41 1.41L9 7.83v5.67C9 16 11 18 13.5 18s4.5-2 4.5-4.5V6h2Z"/>`),
		g.Group(children),
	)
}