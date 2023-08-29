package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolbarOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3ZM5 8h14V5H5v3Zm14 2H5v9h14v-9ZM5 8v2v-2Zm0 0V5v3Zm0 2v9v-9Z"/>`),
		g.Group(children),
	)
}