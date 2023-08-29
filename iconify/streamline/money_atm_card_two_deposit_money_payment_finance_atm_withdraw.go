package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyAtmCardTwoDepositMoneyPaymentFinanceAtmWithdraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1 4.5A.5.5 0 0 1 .5 4V1A.5.5 0 0 1 1 .5h12a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5"/><rect width="7" height="8" x="3.5" y="3" rx=".5"/><circle cx="7" cy="7" r="1.5"/><path d="M3.5 13.5h7"/></g>`),
		g.Group(children),
	)
}