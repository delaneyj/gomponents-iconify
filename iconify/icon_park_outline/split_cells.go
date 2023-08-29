package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitCells(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M4 14V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v38a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9m40 0v9a1 1 0 0 1-1 1H29a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v9M28 24l16 .013m-40 0L20 24"/><path stroke-linejoin="round" d="m39.227 28.778l1.592-1.591L44 24.005l-3.181-3.182l-1.592-1.591M8.755 28.786l-1.59-1.59l-3.183-3.183l3.182-3.182l1.591-1.59"/></g>`),
		g.Group(children),
	)
}