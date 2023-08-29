package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WysiwygSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h14V7H5v12Zm2-7v-2h10v2H7Zm0 4v-2h6v2H7Z"/>`),
		g.Group(children),
	)
}