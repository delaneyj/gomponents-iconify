package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextContinuousTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm5 4a1 1 0 0 1 1-1h11a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1Zm0 4a1 1 0 0 1 1-1h11a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1Zm12 5H4a1 1 0 1 1 0-2h16a1 1 0 1 1 0 2ZM3.293 11.207a1 1 0 0 1 1.414-1.414l1.5 1.5a1 1 0 0 1 0 1.414l-1.5 1.5a1 1 0 0 1-1.414-1.414L4.086 12l-.793-.793Z"/>`),
		g.Group(children),
	)
}