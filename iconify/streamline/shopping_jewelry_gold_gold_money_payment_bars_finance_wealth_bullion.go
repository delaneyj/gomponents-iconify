package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingJewelryGoldGoldMoneyPaymentBarsFinanceWealthBullion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.66 9.3a1 1 0 0 0-1-.8H2.32a1 1 0 0 0-1 .8L.5 13.5h6Zm7 0a1 1 0 0 0-1-.8H9.32a1 1 0 0 0-1 .8l-.82 4.2h6Z"/><path d="m10 8.5l-.84-4.2a1 1 0 0 0-1-.8H5.82a1 1 0 0 0-1 .8L4 8.5"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.68 8.5h4.64"/>`),
		g.Group(children),
	)
}