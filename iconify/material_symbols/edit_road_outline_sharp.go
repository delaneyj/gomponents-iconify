package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditRoadOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11.9V4h2v5.9l-2 2ZM4 20V4h2v16H4Zm6-12V4h2v4h-2Zm0 6v-4h2v4h-2Zm0 6v-4h2v4h-2Zm4 0v-3.175l4.225-4.225l1.075 1.075l-3.8 3.775v1.05h1.05l3.8-3.775l1.05 1.05L17.175 20H14Zm7.4-4.225L18.225 12.6l2.15-2.15l3.175 3.175l-2.15 2.15Z"/>`),
		g.Group(children),
	)
}