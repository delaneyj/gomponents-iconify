package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnBidirectionalDownRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.707 3.293a1 1 0 1 0-1.414 1.414l2.294 2.294H11.5a4.5 4.5 0 0 0-4.5 4.5v7.086l-2.293-2.293a1 1 0 1 0-1.414 1.415l4 4a1 1 0 0 0 1.415 0l4-4a1 1 0 0 0-1.415-1.415L9 18.587v-7.086a2.5 2.5 0 0 1 2.5-2.5h7.086l-2.293 2.293a1 1 0 1 0 1.414 1.414l4-4a1 1 0 0 0 0-1.414l-4-4.001Z"/>`),
		g.Group(children),
	)
}