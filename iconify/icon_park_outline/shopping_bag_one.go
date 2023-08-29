package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="28" height="30" x="10" y="12" rx="3"/><path stroke-linecap="round" d="M30 18v-8a6 6 0 0 0-6-6v0a6 6 0 0 0-6 6v8"/></g>`),
		g.Group(children),
	)
}