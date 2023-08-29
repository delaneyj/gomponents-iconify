package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutputSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v4h-2V5H5v14h14v-2h2v4H3Zm14-4l-1.4-1.4l2.575-2.6H9v-2h9.175L15.6 8.4L17 7l5 5l-5 5Z"/>`),
		g.Group(children),
	)
}