package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FollowUpDateSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5v25.004h36V5M30 37l-6 6l-6-6m6-7v13"/><path stroke-linejoin="round" d="M27.989 10.979L33 16l-5.011 5.09"/><path d="M15 16.001h18"/></g>`),
		g.Group(children),
	)
}