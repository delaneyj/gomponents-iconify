package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExchangeOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M9 18v24h30V18L24 6L9 18Z"/><path d="M17 24h14m-14 6h14m0-6l-5-5m-4 16l-5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExchangeOne0)"/>`),
		g.Group(children),
	)
}