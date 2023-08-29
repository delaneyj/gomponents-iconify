package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksSixSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11V9h3V7H9v10h6v-6h-4Zm0 2h2v2h-2v-2Zm-8 8V3h18v18H3Z"/>`),
		g.Group(children),
	)
}