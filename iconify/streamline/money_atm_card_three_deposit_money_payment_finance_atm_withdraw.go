package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyAtmCardThreeDepositMoneyPaymentFinanceAtmWithdraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="7" x=".5" y="6.5" rx="1"/><path d="M3.5 2v2M7 .5V4m3.5-2v2"/><circle cx="7" cy="10" r="1.5"/></g>`),
		g.Group(children),
	)
}