package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 13v-2h7.6l5-5H14V4h6v6h-2V7.4L12.4 13H4Zm10 7v-2h2.6l-3.2-3.15l1.45-1.45L18 16.6V14h2v6h-6Z"/>`),
		g.Group(children),
	)
}