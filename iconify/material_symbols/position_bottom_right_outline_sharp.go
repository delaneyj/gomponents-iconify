package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PositionBottomRightOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h9v-3H9v3Zm-6 4V3h18v18H3Zm2-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}