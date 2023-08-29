package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomDrawerOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3ZM5 5v8.25L6.25 12h11.5L19 13.25V5H5Zm0 14h14v-3l-2-2H7l-2 2v3Zm0 0h14H5Z"/>`),
		g.Group(children),
	)
}