package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RipplesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h14v-8.55q-.45.275-.963.413T17 11q-1.65 0-2.825-1.175T13 7q0-.525.138-1.038T13.55 5H5v14Zm-2 2V3h18v18H3ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}