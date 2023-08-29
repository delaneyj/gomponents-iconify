package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExchangeThree0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 8.256L24.009 3L42 8.256v10.778A26.316 26.316 0 0 1 24.003 44A26.32 26.32 0 0 1 6 19.029V8.256Z"/><path stroke-linecap="round" d="M17 19h14m-14 6h14m0-6l-5-5m-4 16l-5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExchangeThree0)"/>`),
		g.Group(children),
	)
}