package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TeamDashboardOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm7-2v-6H5v6h5Zm2 0h7v-6h-7v6Zm-7-8h14V5H5v6Z"/>`),
		g.Group(children),
	)
}