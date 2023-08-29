package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveGMobiledataBadgeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V3h22v18H1Zm18-10h-3v2h1v2h-3V9h5V7h-7v10h7v-6ZM5 17h6v-6H7V9h4V7H5v6h4v2H5v2Z"/>`),
		g.Group(children),
	)
}