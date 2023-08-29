package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarketAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 11a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v9h40v-9ZM4.112 39.03l12.176-12.3l6.58 6.3L30.91 26l4.48 4.368"/><path d="M44 18v19a3 3 0 0 1-3 3H12m7.112-26h18M11.11 14h2M4 18v9"/></g>`),
		g.Group(children),
	)
}