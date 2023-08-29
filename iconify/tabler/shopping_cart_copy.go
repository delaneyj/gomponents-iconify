package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartCopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M11.5 17H6V3H4"/><path d="m6 5l14 1l-1 7H6m9 6l2 2l4-4"/></g>`),
		g.Group(children),
	)
}