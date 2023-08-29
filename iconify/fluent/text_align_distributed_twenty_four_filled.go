package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignDistributedTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.293 3.293a1 1 0 0 1 1.414 1.414L5.414 5H21a1 1 0 1 1 0 2H5.414l.293.293a1 1 0 0 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414l2-2ZM18.586 19l-.293.293a1 1 0 0 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414l-2-2a1 1 0 0 0-1.414 1.414l.293.293H3a1 1 0 1 0 0 2h15.586ZM3 11a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H3Z"/>`),
		g.Group(children),
	)
}