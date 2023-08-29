package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopPanelOpenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 16.5l4-4H8l4 4ZM5 19h14v-9H5v9Zm-2 2V3h18v18H3Z"/>`),
		g.Group(children),
	)
}