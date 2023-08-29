package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSnippetSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h12l6 6v12H3Zm4-4h10v-2H7v2Zm0-4h10v-2H7v2Zm0-4h7V7H7v2Z"/>`),
		g.Group(children),
	)
}