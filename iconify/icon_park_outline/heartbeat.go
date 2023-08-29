package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heartbeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke-miterlimit="2" d="m11 32l7-9l6 9l6-9l5 8h9"/><path d="M44 19c0-6.075-4.925-11-11-11c-3.72 0-7.01 1.847-9 4.674A10.987 10.987 0 0 0 15 8C8.925 8 4 12.925 4 19c0 11 13 21 20 23.326c1.194-.397 2.562-1.016 4.01-1.826"/></g>`),
		g.Group(children),
	)
}