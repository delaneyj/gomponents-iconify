package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartCandlestick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 10h-2V6h-2v4h-2v12h2v4h2v-4h2zm-2 10h-2v-8h2zM14 8h-2V4h-2v4H8v10h2v4h2v-4h2zm-2 8h-2v-6h2z"/><path fill="currentColor" d="M30 30H4a2 2 0 0 1-2-2V2h2v26h26Z"/>`),
		g.Group(children),
	)
}