package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardHideSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l-4-4h8l-4 4ZM2 17V3h20v14H2Zm6-3h8v-2H8v2Zm-3-3h2V9H5v2Zm3 0h2V9H8v2Zm3 0h2V9h-2v2Zm3 0h2V9h-2v2Zm3 0h2V9h-2v2ZM5 8h2V6H5v2Zm3 0h2V6H8v2Zm3 0h2V6h-2v2Zm3 0h2V6h-2v2Zm3 0h2V6h-2v2Z"/>`),
		g.Group(children),
	)
}