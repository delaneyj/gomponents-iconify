package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCurrencyBitcoinCryptoCirclePaymentBlokchainFinanceBitcoinCurrencyMoney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.79 6.45h0A2.23 2.23 0 0 0 7.83 2H4.39a.81.81 0 0 0-.81.81v8.09a.81.81 0 0 0 .81.81h3.4a2.63 2.63 0 0 0 0-5.26ZM5.13 2V.5M7.63 2V.5m-2.5 13V12m2.5 1.5V12m.2-5.55H3.58"/>`),
		g.Group(children),
	)
}