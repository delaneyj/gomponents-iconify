package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCashBillTwoCurrencyBillingPaymentFinanceCashBillMoneyAccounting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2.5" rx="1"/><circle cx="7" cy="7" r="1.5"/><path d="M3 5h.5m7 4h.5"/></g>`),
		g.Group(children),
	)
}