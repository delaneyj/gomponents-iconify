package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipExitSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-9h9V4h11v16H2Zm15.075-3.5l1.425-1.425L15.4 12H18v-2h-6v6h2v-2.575l3.075 3.075ZM2 9V4h7v5H2Z"/>`),
		g.Group(children),
	)
}