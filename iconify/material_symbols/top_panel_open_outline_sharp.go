package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopPanelOpenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 16.5l4-4H8l4 4ZM5 8h14V5H5v3Zm0 11h14v-9H5v9ZM5 8V5v3ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}