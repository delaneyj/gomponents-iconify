package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesDiamondDiamondMoneyPaymentFinanceWealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.89.93a1 1 0 0 0-.82-.43H3.93a1 1 0 0 0-.82.43l-1.6 2.31a1 1 0 0 0 .15 1.28l4.67 4.72a1 1 0 0 0 1.34 0l4.67-4.72a1 1 0 0 0 .15-1.28Z"/><path d="M7 .5L4.68 3.94l2.18 5.55M7 .5l2.32 3.44l-2.18 5.55M1.35 3.94h11.3M13 13.5h0a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2h0"/></g>`),
		g.Group(children),
	)
}