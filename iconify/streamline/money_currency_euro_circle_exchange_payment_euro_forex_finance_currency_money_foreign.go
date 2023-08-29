package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyEuroCircleExchangePaymentEuroForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 9.8a2.59 2.59 0 0 1-1 .2a2.88 2.88 0 0 1-2.75-3A2.88 2.88 0 0 1 7.5 4a2.76 2.76 0 0 1 .82.13M3.5 6h3m-3 2h3"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}