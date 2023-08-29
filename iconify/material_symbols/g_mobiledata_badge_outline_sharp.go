package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GMobiledataBadgeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17h8v-6h-4v2h2v2h-4V9h6V7H8v10Zm-5 4V3h18v18H3Zm2-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}