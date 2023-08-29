package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarketAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMarketAnalysis0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M44 11a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v9h40v-9Z"/><path stroke="#fff" d="m4.112 39.03l12.176-12.3l6.58 6.3L30.91 26l4.48 4.368"/><path stroke="#fff" d="M44 18v19a3 3 0 0 1-3 3H12"/><path stroke="#000" d="M19.112 14h18M11.11 14h2"/><path stroke="#fff" d="M4 18v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMarketAnalysis0)"/>`),
		g.Group(children),
	)
}