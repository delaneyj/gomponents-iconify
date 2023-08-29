package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodToastBreadToastBreakfast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 3A2.5 2.5 0 0 0 11 .5H3A2.5 2.5 0 0 0 1.5 5v7.5a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V5a2.49 2.49 0 0 0 1-2Zm-5 .5l-4 4m5-1l-4 4"/>`),
		g.Group(children),
	)
}