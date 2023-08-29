package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9h2V7H7v2Zm4 0h2V7h-2v2Zm4 0h2V7h-2v2ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}