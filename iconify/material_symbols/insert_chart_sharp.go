package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertChartSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h2v-7H7v7Zm4 0h2V7h-2v10Zm4 0h2v-4h-2v4ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}