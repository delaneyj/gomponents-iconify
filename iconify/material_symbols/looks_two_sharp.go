package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h6v-2h-4v-2h4V7H9v2h4v2H9v6Zm-6 4V3h18v18H3Z"/>`),
		g.Group(children),
	)
}