package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatasetOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Zm2-8h4V7H7v4Zm6 0h4V7h-4v4Zm-6 6h4v-4H7v4Zm6 0h4v-4h-4v4Z"/>`),
		g.Group(children),
	)
}