package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieEditOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19V3h2l2 4h3L7 3h2l2 4h3l-2-4h2l2 4h3l-2-4h5v6H4v8h8v2H2Zm16.3-6.475l1.075 1.075l-3.875 3.85v1.05h1.05l3.875-3.85l1.05 1.05l-4.3 4.3H14v-3.175l4.3-4.3Zm3.175 3.175L18.3 12.525l2.15-2.15l3.175 3.175l-2.15 2.15Z"/>`),
		g.Group(children),
	)
}