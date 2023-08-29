package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkingWaterFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6 1a2 2 0 0 0-2 2v3.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 1 .5-.5H14V1H6z" fill="currentColor"/><path d="M7 15H4a.5.5 0 0 1-.48-.38L2 8.62A.5.5 0 0 1 2.5 8h6a.5.5 0 0 1 .5.62l-1.5 6A.5.5 0 0 1 7 15zm-3.35-4h3.71l.5-2H3.14z" fill="currentColor"/>`),
		g.Group(children),
	)
}