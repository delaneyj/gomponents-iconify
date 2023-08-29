package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H3V3h18v18Zm-8-8v6h6v-6h-6Zm0-2h6V5h-6v6Zm-2 0V5H5v6h6Zm0 2H5v6h6v-6Z"/>`),
		g.Group(children),
	)
}