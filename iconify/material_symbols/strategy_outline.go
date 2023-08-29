package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrategyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 11L2 9V5l3.5-2L9 5v4l-3.5 2Zm0-2.3L7 7.85v-1.7L5.5 5.3L4 6.15v1.7l1.5.85Zm11 3.075V9.45L20 11.5v7L14 22l-6-3.5v-7l3.5-2.05v2.325l-1.5.875v4.7l4 2.325l4-2.325v-4.7l-1.5-.875ZM13 14V2h9l-2 3l2 3h-7v6h-2Zm1 1.725ZM5.5 7Z"/>`),
		g.Group(children),
	)
}