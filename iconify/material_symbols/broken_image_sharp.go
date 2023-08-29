package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrokenImageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-8.6l3 3l4-4l4 4l4-4l3 3V21H3ZM3 3h18v8.575l-3-3l-4 4l-4-4l-4 4l-3-3V3Z"/>`),
		g.Group(children),
	)
}