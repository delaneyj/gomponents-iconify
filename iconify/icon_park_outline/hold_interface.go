package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoldInterface(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m20 33l6 2s15-3 17-3s2 2 0 4s-9 8-15 8s-10-3-14-3H4"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 29c2-2 6-5 10-5s13.5 4 15 6s-3 5-3 5"/><rect width="16" height="6" x="26" y="15" rx="3"/><path stroke-linecap="round" d="M26 9h16"/></g>`),
		g.Group(children),
	)
}