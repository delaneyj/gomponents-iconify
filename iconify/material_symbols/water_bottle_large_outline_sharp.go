package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterBottleLargeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20h10v-2h-2v-7h-2V9h4V7H7v2h2v7h2v2H7v2Zm-2 2v-6h2v-5H5V5h5V4H9V2h6v2h-1v1h5v6h-2v5h2v6H5Zm7-8.5Z"/>`),
		g.Group(children),
	)
}