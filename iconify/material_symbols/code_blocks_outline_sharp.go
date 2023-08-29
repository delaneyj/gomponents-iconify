package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeBlocksOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.6 15.6l1.4-1.425L8.825 12L11 9.825L9.6 8.4L6 12l3.6 3.6Zm4.8 0L18 12l-3.6-3.6L13 9.825L15.175 12L13 14.175l1.4 1.425ZM3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}