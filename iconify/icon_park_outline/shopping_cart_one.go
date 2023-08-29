package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="20.5" cy="41.5" r="3.5" fill="currentColor"/><circle cx="37.5" cy="41.5" r="3.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m5 6l9 6l5 22h20l5-17H25m0 9h16"/></g>`),
		g.Group(children),
	)
}