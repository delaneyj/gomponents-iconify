package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h14v-7H5v7Zm0-9h14V5H5v5Zm7-2h6V6h-6v2Zm-7 2V5v5Z"/>`),
		g.Group(children),
	)
}