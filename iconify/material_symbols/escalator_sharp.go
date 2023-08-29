package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 18h4.8l5-9h3.2V6h-4.8l-5 9H5.5v3ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}