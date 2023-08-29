package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCashBagBagPaymentCashMoneyFinance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 3.02L13.5.5M8 3.02l5.5 1.48m-7 9c3.5 0 6-1.24 6-4c0-3-1.5-5-4.5-6.5l1.18-1.53a.65.65 0 0 0-.56-.95H4.38a.65.65 0 0 0-.56 1L5 3C2 4.52.5 6.52.5 9.52c0 2.74 2.5 3.98 6 3.98Z"/>`),
		g.Group(children),
	)
}