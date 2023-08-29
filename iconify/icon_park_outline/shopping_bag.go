package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 17h38l-4.2 26H9.2L5 17Zm30 0c0-6.627-4.925-12-11-12s-11 5.373-11 12"/><circle cx="17" cy="26" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 33s2 3 6 3s6-3 6-3"/><circle cx="31" cy="26" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}