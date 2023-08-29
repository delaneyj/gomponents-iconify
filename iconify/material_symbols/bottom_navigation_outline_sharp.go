package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomNavigationOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-6h14V5H5v10Zm0 2v2h14v-2H5Zm0 0v2v-2Z"/>`),
		g.Group(children),
	)
}