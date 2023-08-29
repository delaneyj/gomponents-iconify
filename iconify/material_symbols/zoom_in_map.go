package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomInMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.9 21.5l-1.4-1.4L5.6 17H3v-2h6v6H7v-2.6l-3.1 3.1Zm16.2 0L17 18.4V21h-2v-6h6v2h-2.6l3.1 3.1l-1.4 1.4ZM3 9V7h2.6L2.5 3.9l1.4-1.4L7 5.6V3h2v6H3Zm12 0V3h2v2.6l3.1-3.1l1.4 1.4L18.4 7H21v2h-6Z"/>`),
		g.Group(children),
	)
}