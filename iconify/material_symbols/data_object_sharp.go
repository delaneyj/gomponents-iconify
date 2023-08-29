package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataObjectSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20v-2h4v-5.675h2v-.65h-2V6h-4V4h6v5.85h2v4.3h-2V20h-6ZM4 20v-5.85H2v-4.3h2V4h6v2H6v5.675H4v.65h2V18h4v2H4Z"/>`),
		g.Group(children),
	)
}