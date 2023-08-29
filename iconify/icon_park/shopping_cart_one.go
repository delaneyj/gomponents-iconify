package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="20.5" cy="41.5" r="3.5" fill="#000"/><circle cx="37.5" cy="41.5" r="3.5" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 6L14 12L19 34H39L44 17H25"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M25 26L32.2727 26L41 26"/></g>`),
		g.Group(children),
	)
}