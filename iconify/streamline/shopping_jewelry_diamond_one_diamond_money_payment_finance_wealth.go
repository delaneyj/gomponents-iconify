package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingJewelryDiamondOneDiamondMoneyPaymentFinanceWealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 5.5l-6.5 7l-6.5-7l3-4h7l3 4zm-13 0h13"/><path d="m3.5 1.5l3.5 4l3.5-4M7 5.5v7"/></g>`),
		g.Group(children),
	)
}