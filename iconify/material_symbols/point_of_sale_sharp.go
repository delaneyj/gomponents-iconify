package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointOfSaleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 8V2h14v6H5Zm2-2h10V4H7v2ZM2 22v-3h20v3H2Zm0-4l4-9h12l4 9H2Zm6-2h2v-1H8v1Zm0-2h2v-1H8v1Zm0-2h2v-1H8v1Zm3 4h2v-1h-2v1Zm0-2h2v-1h-2v1Zm0-2h2v-1h-2v1Zm3 4h2v-1h-2v1Zm0-2h2v-1h-2v1Zm0-2h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}