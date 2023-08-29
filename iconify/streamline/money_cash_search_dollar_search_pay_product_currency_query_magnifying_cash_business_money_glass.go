package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCashSearchDollarSearchPayProductCurrencyQueryMagnifyingCashBusinessMoneyGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.5" cy="6.5" r="6"/><path d="m10.74 10.74l2.76 2.76M6.5 4V2.5M5 8c0 .75.67 1 1.5 1S8 9 8 8c0-1.5-3-1.5-3-3c0-1 .67-1 1.5-1S8 4.38 8 5M6.5 9v1.5"/></g>`),
		g.Group(children),
	)
}