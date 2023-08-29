package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionAdjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m20 8l4-4m0 0l4 4m-4-4v12m-4 24l4 4m0 0l4-4m-4 4V32m16-12l4 4m0 0l-4 4m4-4H32M8 20l-4 4m0 0l4 4m-4-4h12"/><circle cx="24" cy="24" r="2"/></g>`),
		g.Group(children),
	)
}