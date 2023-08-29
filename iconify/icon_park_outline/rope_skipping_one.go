package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RopeSkippingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke-miterlimit="2" d="M7 40V22S8 8 15 5s15 6 15 15c0 5-2 9-6 9s-6-4-6-9c0-9 8-18 15-15s8 17 8 17v18"/><path d="M4 31h6m28 0h6"/></g>`),
		g.Group(children),
	)
}