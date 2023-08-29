package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StockMarket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTStockMarket0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 20h8v14H6zm14-6h8v26h-8z"/><path stroke-linecap="round" d="M24 44v-4"/><path fill="#555" d="M34 12h8v9h-8z"/><path stroke-linecap="round" d="M10 20V10m28 24V21m0-9V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTStockMarket0)"/>`),
		g.Group(children),
	)
}