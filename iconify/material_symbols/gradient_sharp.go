package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradientSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h-2v-2h2v-2h-2v-2h-2v2h-2v-2h-2v2H9v-2H7v2H5v2h2v2H5v2Zm2-4h2v-2H7v2Zm2 2h2v-2H9v2Zm2-2h2v-2h-2v2Zm2 2h2v-2h-2v2Zm2-2h2v-2h-2v2Zm4 4v-2v2Z"/>`),
		g.Group(children),
	)
}