package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M12 17H6V3H4"/><path d="m6 5l14 1l-.716 5.011M14 13H6m15.121 7.121a3 3 0 1 0-4.242 0c.418.419 1.125 1.045 2.121 1.879c1.051-.89 1.759-1.516 2.121-1.879zM19 18v.01"/></g>`),
		g.Group(children),
	)
}