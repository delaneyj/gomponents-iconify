package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 3a1 1 0 1 1 0 2H6.414l14.293 14.293a1 1 0 0 1-1.414 1.414L5 6.414V13a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1h9Z"/>`),
		g.Group(children),
	)
}