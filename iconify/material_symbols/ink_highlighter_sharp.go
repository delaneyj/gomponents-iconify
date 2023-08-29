package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkHighlighterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1.5 21l3.15-3.15l-.75-.75v-1.4L9.6 10l5.4 5.4l-5.7 5.7H7.9l-.75-.75l-.65.65h-5Zm9.525-12.425L17.4 2.2l5.4 5.4l-6.375 6.375l-5.4-5.4Z"/>`),
		g.Group(children),
	)
}