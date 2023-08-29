package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 21a1 1 0 0 0 0-2H6.414L20.707 4.707a1 1 0 0 0-1.414-1.414L5 17.586V11a1 1 0 1 0-2 0v9a1 1 0 0 0 1 1h9Z"/>`),
		g.Group(children),
	)
}