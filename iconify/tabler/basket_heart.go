package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 10l-2-6m-8 6l2-6m1.5 16H7.244a3 3 0 0 1-2.965-2.544l-1.255-7.152A2 2 0 0 1 5.001 8H19a2 2 0 0 1 1.977 2.304l-.143.817"/><path d="M12.602 12.092a2 2 0 0 0-2.233 3.066M18 22l3.35-3.284a2.143 2.143 0 0 0 .005-3.071a2.242 2.242 0 0 0-3.129-.006l-.224.22l-.223-.22a2.242 2.242 0 0 0-3.128-.006a2.143 2.143 0 0 0-.006 3.071L18 22z"/></g>`),
		g.Group(children),
	)
}