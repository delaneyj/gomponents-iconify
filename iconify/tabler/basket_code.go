package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 10l-2-6m-8 6l2-6m2 16H7.244a3 3 0 0 1-2.965-2.544l-1.255-7.152A2 2 0 0 1 5.001 8H19a2 2 0 0 1 1.977 2.304c-.21 1.202-.37 2.104-.475 2.705"/><path d="M10 14a2 2 0 1 0 4 0a2 2 0 0 0-4 0m10 7l2-2l-2-2m-3 0l-2 2l2 2"/></g>`),
		g.Group(children),
	)
}