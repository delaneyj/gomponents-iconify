package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHangingTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 5a1 1 0 1 1 0 2H3a1 1 0 0 1 0-2h18Zm-7 12a1 1 0 1 1 0 2H3a1 1 0 1 1 0-2h11Zm1-5a1 1 0 0 0-1-1H3a1 1 0 1 0 0 2h11a1 1 0 0 0 1-1Zm4.293 1.293a1 1 0 0 1 1.414 1.414L19.414 16l1.293 1.293a1 1 0 0 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414l2-2Z"/>`),
		g.Group(children),
	)
}