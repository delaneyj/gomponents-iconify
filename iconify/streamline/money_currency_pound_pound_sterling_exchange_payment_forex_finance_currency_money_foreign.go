package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyPoundPoundSterlingExchangePaymentForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.75 4.75V4A3.26 3.26 0 0 0 8.5.75h0A3.26 3.26 0 0 0 5.25 4v6.25l-2 3h8.5m-9.5-6h5.99"/>`),
		g.Group(children),
	)
}