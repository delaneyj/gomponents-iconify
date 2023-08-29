package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatShapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 23v-6h2V7H1V1h6v2h10V1h6v6h-2v10h2v6h-6v-2H7v2H1Zm6-4h10v-2h2V7h-2V5H7v2H5v10h2v2Zm.8-3l3.4-9h1.6l3.4 9h-1.55l-.8-2.3H10.2L9.4 16H7.8Zm2.85-3.6h2.7l-1.3-3.75h-.1l-1.3 3.75Z"/>`),
		g.Group(children),
	)
}