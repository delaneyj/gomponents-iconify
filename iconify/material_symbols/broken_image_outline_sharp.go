package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrokenImageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm3-8.425l4-4l4 4l4-4l1 1V5H5v6.575l1 1ZM5 19h14v-6.6l-1-1l-4 4l-4-4l-4 4l-1-1V19Zm0 0v-6.6v2V5v14Z"/>`),
		g.Group(children),
	)
}