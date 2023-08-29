package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvMobiledataBadgeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17h6v-2H7v-2h3v-2H7V9h4V7H5v10Zm9.5 0h2L19 7h-2l-1.5 6L14 7h-2l2.5 10ZM1 21V3h22v18H1Zm2-2h18V5H3v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}