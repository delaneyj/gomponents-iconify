package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M8 7h32v24H8z"/><path stroke-linecap="round" d="M4 7h40M15 41l9-10l9 10M16 13h16m-16 6h12m-12 6h6"/></g>`),
		g.Group(children),
	)
}