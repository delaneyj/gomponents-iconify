package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropdownOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13h7V6h-7v7Zm-8 8V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}