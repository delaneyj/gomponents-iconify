package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddToHomeScreenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23v-6h2v1h10V6H8v1H6V1h14v22H6Zm-1.6-6L3 15.6L8.6 10H5V8h7v7h-2v-3.6L4.4 17Z"/>`),
		g.Group(children),
	)
}