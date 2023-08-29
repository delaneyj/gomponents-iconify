package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18h11v-2.675H9V18ZM4 8.675h3V6H4v2.675Zm0 4.675h3v-2.675H4v2.675ZM4 18h3v-2.675H4V18Zm5-4.65h11v-2.675H9v2.675Zm0-4.675h11V6H9v2.675ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}