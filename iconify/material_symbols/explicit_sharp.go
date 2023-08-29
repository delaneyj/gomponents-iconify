package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExplicitSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h6v-2h-4v-2h4v-2h-4V9h4V7H9v10Zm-6 4V3h18v18H3Z"/>`),
		g.Group(children),
	)
}