package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFirstLineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.293 3.293a1 1 0 1 1 1.414 1.414L20.414 6l1.293 1.293a1 1 0 0 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414l2-2ZM14 5a1 1 0 1 1 0 2H3a1 1 0 0 1 0-2h11ZM3 17h18a1 1 0 1 1 0 2H3a1 1 0 1 1 0-2Zm19-5a1 1 0 0 0-1-1H3a1 1 0 1 0 0 2h18a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}