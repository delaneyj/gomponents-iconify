package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddToHomeScreenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23v-6h2v1h10V6H8v1H6V1h14v22H6Zm2-3v1h10v-1H8Zm-3.6-3L3 15.6L8.6 10H5V8h7v7h-2v-3.6L4.4 17ZM8 4h10V3H8v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}