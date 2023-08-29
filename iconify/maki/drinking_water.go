package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkingWater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 1a2 2 0 0 0-2 2v3.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 1 .5-.5H14V1Zm1 14H4a.5.5 0 0 1-.48-.38L2 8.62a.5.5 0 0 1 .365-.606A.558.558 0 0 1 2.5 8h6a.5.5 0 0 1 .514.485A.47.47 0 0 1 9 8.62l-1.5 6A.5.5 0 0 1 7 15Zm-3.35-4h3.71l.5-2H3.14Z"/>`),
		g.Group(children),
	)
}