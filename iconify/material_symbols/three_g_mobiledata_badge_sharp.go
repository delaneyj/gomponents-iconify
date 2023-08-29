package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeGMobiledataBadgeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17h7v-6h-3v2h1v2h-3V9h5V7h-7v10Zm-7 0h6v-4l-1-1l1-1V7H5v2h4v2H5v2h4v2H5v2Zm-4 4V3h22v18H1Z"/>`),
		g.Group(children),
	)
}