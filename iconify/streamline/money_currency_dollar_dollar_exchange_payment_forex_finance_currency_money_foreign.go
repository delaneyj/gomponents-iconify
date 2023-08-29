package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyDollarDollarExchangePaymentForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 10.5h3.85a1.65 1.65 0 0 0 .32-3.27l-3.34-.46a1.65 1.65 0 0 1 .32-3.27H9.5M7 13.5V.5"/>`),
		g.Group(children),
	)
}