package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanToolAltOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.525 21L1.15 11.925l1.775-1.7L7 13.075V2h2v14.925l-2.775-1.95L9.475 19H19V9h2v12H8.525ZM11 13V6h2v7h-2Zm4 0V7h2v6h-2Zm-1 2Z"/>`),
		g.Group(children),
	)
}