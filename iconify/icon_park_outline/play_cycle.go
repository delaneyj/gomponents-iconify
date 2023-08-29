package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 25c0-6.65 5.396-12 12-12h28"/><path d="m38 7l6 6l-6 6m6 4c0 6.65-5.396 12-12 12H4"/><path d="m10 41l-6-6l6-6"/></g>`),
		g.Group(children),
	)
}