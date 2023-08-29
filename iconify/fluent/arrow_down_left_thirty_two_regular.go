package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.004 28.996a1 1 0 1 0 0-2H6.418L28.707 4.707a1 1 0 0 0-1.414-1.414L5.003 25.582V14.996a1 1 0 1 0-2 0v13a1 1 0 0 0 1 1h13Z"/>`),
		g.Group(children),
	)
}