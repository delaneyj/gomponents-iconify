package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopPanelCloseOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16.5h8l-4-4l-4 4ZM5 8h14V5H5v3Zm0 11h14v-9H5v9ZM5 8V5v3ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}