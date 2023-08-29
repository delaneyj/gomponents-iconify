package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterBottleLargeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22v-6h2v-5H5V5h5V4H9V2h6v2h-1v1h5v6h-2v5h2v6H5Z"/>`),
		g.Group(children),
	)
}