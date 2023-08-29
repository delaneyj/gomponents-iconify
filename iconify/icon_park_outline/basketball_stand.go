package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketballStand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="5" rx="2"/><path d="M31 22v-7H17v7m1 17h12m-13-6h14m1-6l-2 12l1 4M16 27l2 12l-1 4m7-16v16m10-16H14"/></g>`),
		g.Group(children),
	)
}