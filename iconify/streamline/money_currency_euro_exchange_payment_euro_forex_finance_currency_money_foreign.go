package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyEuroExchangePaymentEuroForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 12.59a5 5 0 0 1-2 .41a5.77 5.77 0 0 1-5.5-6A5.77 5.77 0 0 1 10 1a4.89 4.89 0 0 1 1.63.27M2 5.5h6m-6 3h6"/>`),
		g.Group(children),
	)
}