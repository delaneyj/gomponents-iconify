package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="4" d="M5.939 13.934L23.036 4.53a2 2 0 0 1 1.928 0l17.097 9.404a2 2 0 0 1 .683 2.888l-17.098 24.79a2 2 0 0 1-3.292 0L5.256 16.823a2 2 0 0 1 .683-2.888Z"/>`),
		g.Group(children),
	)
}