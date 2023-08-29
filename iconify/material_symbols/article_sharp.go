package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h7v-2H7v2Zm0-4h10v-2H7v2Zm0-4h10V7H7v2ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}