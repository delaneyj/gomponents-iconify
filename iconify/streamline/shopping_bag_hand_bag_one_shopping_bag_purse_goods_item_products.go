package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagHandBagOneShoppingBagPurseGoodsItemProducts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.27 13.5H2.73a2 2 0 0 1-2-2.22l.67-5.89a1 1 0 0 1 1-.89h9.2a1 1 0 0 1 1 .89l.65 5.89a2 2 0 0 1-1.98 2.22Z"/><path d="M3 4.5a4 4 0 0 1 8 0m-6.5 3h5"/></g>`),
		g.Group(children),
	)
}