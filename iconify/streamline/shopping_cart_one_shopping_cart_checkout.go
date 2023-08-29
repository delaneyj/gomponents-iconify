package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartOneShoppingCartCheckout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.5H11l-.87 8.65a1 1 0 0 1-1 .85h-6.3a1 1 0 0 1-1-.68l-1.33-4a1 1 0 0 1 .14-.9A1 1 0 0 1 1.5 4h9.15"/><circle cx="3" cy="13" r=".5"/><circle cx="9.5" cy="13" r=".5"/></g>`),
		g.Group(children),
	)
}