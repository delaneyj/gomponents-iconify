package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="28" r="16"/><path stroke-linecap="round" stroke-linejoin="round" d="M28 4h-8m4 0v8m11 4l3-3M24 28v-6m0 6h-6"/></g>`),
		g.Group(children),
	)
}