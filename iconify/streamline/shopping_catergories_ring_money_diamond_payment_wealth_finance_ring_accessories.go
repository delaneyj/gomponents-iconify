package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCatergoriesRingMoneyDiamondPaymentWealthFinanceRingAccessories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="9" r="4.5"/><path d="m8.5.5l1 1l-2.5 3l-2.5-3l1-1h3z"/></g>`),
		g.Group(children),
	)
}