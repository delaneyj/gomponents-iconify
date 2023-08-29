package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradientOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13v-2h2v2h-2Zm-2 2v-2h2v2H9Zm4 0v-2h2v2h-2Zm2-2v-2h2v2h-2Zm-8 0v-2h2v2H7Zm-4 8V3h18v18H3Zm4-2h2v-2H7v2Zm4 0h2v-2h-2v2Zm8 0v-2v2ZM5 17h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h-2v-2h2V5H5v8h2v2H5v2Zm0 2V5v14Zm14-6v2v-2Zm-4 4v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}