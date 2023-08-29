package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M12 17H6V3H4"/><path d="m6 5l14 1l-.857 5.998M15.5 13H6m10 6a3 3 0 1 0 6 0a3 3 0 1 0-6 0m1 2l4-4"/></g>`),
		g.Group(children),
	)
}