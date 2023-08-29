package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarArrowCounterclockwiseThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 9a5 5 0 0 1 5-5h14a5 5 0 0 1 5 5v14a5 5 0 0 1-5 5H9a5 5 0 0 1-5-5v-4a1 1 0 1 1 2 0v4a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V12H10.293a2.508 2.508 0 0 0 0-2H26V9a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v2.586l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L4 11.586V9Zm6.5 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm4 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5-6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm4 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}