package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 14a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M5.001 8H19a2 2 0 0 1 1.977 2.304l-1.255 7.152A3 3 0 0 1 16.756 20H7.244a3 3 0 0 1-2.965-2.544l-1.255-7.152A2 2 0 0 1 5.001 8zM17 10l-2-6m-8 6l2-6"/></g>`),
		g.Group(children),
	)
}