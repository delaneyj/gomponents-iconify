package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrapTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1Zm1 11a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H3Zm12.414 0H19a2 2 0 1 0 0-4H3a1 1 0 1 1 0-2h16a4 4 0 0 1 0 8h-3.586l.293.293a1 1 0 0 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414l2-2a1 1 0 0 1 1.414 1.414l-.293.293Z"/>`),
		g.Group(children),
	)
}