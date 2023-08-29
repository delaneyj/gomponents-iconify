package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoStableSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.975 18l2.3-8.65L7.075 6l-2.3 8.65l12.2 3.35ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}