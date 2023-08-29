package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabMoveSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-5h2v3h14V7H5v3H3V3h18v18H3Zm8.5-3.5l-1.4-1.4l2.075-2.1H3v-2h9.175L10.1 9.9l1.4-1.4L16 13l-4.5 4.5Z"/>`),
		g.Group(children),
	)
}