package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalSpacingBetweenItems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M8 6v6h32V6"/><path d="M14 24h20"/><path stroke-linejoin="round" d="M8 42v-6h32v6"/></g>`),
		g.Group(children),
	)
}