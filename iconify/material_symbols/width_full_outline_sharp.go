package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidthFullOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h1V6H4v12Zm3 0h10V6H7v12Zm12 0h1V6h-1v12ZM7 6v12V6Z"/>`),
		g.Group(children),
	)
}