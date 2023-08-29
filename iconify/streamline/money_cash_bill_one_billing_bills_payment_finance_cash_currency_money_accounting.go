package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCashBillOneBillingBillsPaymentFinanceCashCurrencyMoneyAccounting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y=".5" rx="1"/><circle cx="7" cy="4.5" r="1.5"/><path d="M1.5 11h11m-10 2.5h9"/></g>`),
		g.Group(children),
	)
}