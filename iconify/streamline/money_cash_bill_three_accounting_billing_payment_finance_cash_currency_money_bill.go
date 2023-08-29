package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCashBillThreeAccountingBillingPaymentFinanceCashCurrencyMoneyBill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10.5" height="8" x=".5" y="1.75" rx="1"/><circle cx="5.75" cy="5.75" r="1.5"/><path d="M3.5 12.25h9a1 1 0 0 0 1-1v-5"/></g>`),
		g.Group(children),
	)
}