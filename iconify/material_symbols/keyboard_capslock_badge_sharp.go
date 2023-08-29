package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardCapslockBadgeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h10v-2H7v2Zm1.4-4L12 9.4l3.6 3.6l1.4-1.4l-5-5l-5 5L8.4 13ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}