package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextCollapseTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 4a1 1 0 0 0 1 1h14a1 1 0 1 0 0-2H7a1 1 0 0 0-1 1Zm7 5a1 1 0 0 0 1 1h7a1 1 0 1 0 0-2h-7a1 1 0 0 0-1 1Zm-6 9h14a1 1 0 1 1 0 2H7a1 1 0 1 1 0-2Zm6-4a1 1 0 0 0 1 1h7a1 1 0 1 0 0-2h-7a1 1 0 0 0-1 1Zm-2-2.5a4.5 4.5 0 1 0-9 0a4.5 4.5 0 0 0 9 0Zm-2 0a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1 0-1h4a.5.5 0 0 1 .5.5Z"/>`),
		g.Group(children),
	)
}