package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnDownRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.72 16.78a.75.75 0 0 0 1.06 0l4-4a.75.75 0 0 0 0-1.06l-4-4a.75.75 0 1 0-1.06 1.06l2.72 2.72H7A1.5 1.5 0 0 1 5.5 10V3.75a.75.75 0 0 0-1.5 0V10a3 3 0 0 0 3 3h6.44l-2.72 2.72a.75.75 0 0 0 0 1.06Z"/>`),
		g.Group(children),
	)
}