package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookArrowClockwiseTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 4.25a.75.75 0 0 1-.75.75H18.5a.75.75 0 0 1 0-1.5h1.33l-.008-.006l-.01-.007A4.5 4.5 0 1 0 21.5 7A.75.75 0 0 1 23 7a6 6 0 1 1-2.5-4.874V1.5a.746.746 0 0 1 .75-.75a.75.75 0 0 1 .75.75v2.75ZM12.101 2H6.5A2.5 2.5 0 0 0 4 4.5v15A2.5 2.5 0 0 0 6.5 22h13.25a.75.75 0 0 0 0-1.5H6.5a1 1 0 0 1-1-1h14.25a.75.75 0 0 0 .75-.75v-5.687A7.003 7.003 0 0 1 10.07 8a7.062 7.062 0 0 1 .22-3a7.018 7.018 0 0 1 1.81-3Z"/>`),
		g.Group(children),
	)
}