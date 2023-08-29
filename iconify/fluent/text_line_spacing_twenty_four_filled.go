package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextLineSpacingTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 6.414V10a1 1 0 1 1-2 0V6.414l-.293.293a1 1 0 1 1-1.414-1.414l2-2a.998.998 0 0 1 1.414 0l2 2a1 1 0 0 1-1.414 1.414L20 6.414ZM2 6a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1Zm0 6a1 1 0 0 1 1-1h11a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1Zm1 5a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H3Zm17-3v3.586l.293-.293a1 1 0 0 1 1.414 1.414l-2 2a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414l.293.293V14a1 1 0 1 1 2 0Z"/>`),
		g.Group(children),
	)
}