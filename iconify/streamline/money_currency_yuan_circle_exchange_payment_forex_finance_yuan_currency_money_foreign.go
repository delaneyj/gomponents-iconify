package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyYuanCircleExchangePaymentForexFinanceYuanCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m5 4l2.08 2.77L9.15 4M7.08 6.77V10m-1.62-.69h3.23M5.46 7.5h3.23"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}