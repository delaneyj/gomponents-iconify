package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyPoundCirclePoundSterlingExchangePaymentForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.06 5.42v-.36A1.56 1.56 0 0 0 7.5 3.5h0a1.56 1.56 0 0 0-1.56 1.56v3L5 9.5h4.06M4.5 6.62h2.88"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}