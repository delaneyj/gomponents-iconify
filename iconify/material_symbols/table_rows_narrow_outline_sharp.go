package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRowsNarrowOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15h14v-2H5v2Zm0-4h14V9H5v2Zm0-4h14V5H5v2ZM3 21V3h18v18H3Zm2-2h14v-2H5v2Z"/>`),
		g.Group(children),
	)
}