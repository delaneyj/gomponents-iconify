package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyYuanExchangePaymentForexFinanceYuanCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5.5l4.5 6l4.5-6M7 6.5v7m-3.5-3h7m-7-2h7"/>`),
		g.Group(children),
	)
}