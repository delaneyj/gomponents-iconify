package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartBasketThreeShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.36 6.62a1 1 0 0 0-.24-.78a1 1 0 0 0-.75-.34H1.63a1 1 0 0 0-.75.34a1 1 0 0 0-.24.78l.75 6a1 1 0 0 0 1 .88h9.24a1 1 0 0 0 1-.88ZM2.5 5.5V5a4.5 4.5 0 0 1 9 0v.5M5 8.5v2m4-2v2"/>`),
		g.Group(children),
	)
}