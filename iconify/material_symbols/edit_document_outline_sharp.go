package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditDocumentOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h10l6 6v4h-2V9h-5V4H6v16h6v2H4Zm2-2V4v16Zm12.275-5.45l1.075 1.05l-3.85 3.85v1.05h1.05l3.85-3.85l1.075 1.05l-4.3 4.3H14v-3.175l4.275-4.275Zm3.2 3.15l-3.2-3.15l2.15-2.15l3.175 3.175l-2.125 2.125Z"/>`),
		g.Group(children),
	)
}