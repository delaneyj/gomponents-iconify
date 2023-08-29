package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkEraserOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.25 18H22v2h-6.75l2-2Zm-12.5 2L1.2 16.45L15 2.15l7.8 7.8L13 20H4.75Zm7.4-2L20 9.95L15.05 5L4 16.4L5.6 18h6.55ZM12 12Z"/>`),
		g.Group(children),
	)
}