package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StockMarket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 20h8v14H6zm14-6h8v26h-8z"/><path stroke-linecap="round" d="M24 44v-4"/><path d="M34 12h8v9h-8z"/><path stroke-linecap="round" d="M10 20V10m28 24V21m0-9V4"/></g>`),
		g.Group(children),
	)
}