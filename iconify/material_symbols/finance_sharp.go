package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinanceSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h2v16h16v2H3Zm3-3V9h4v9H6Zm5 0V4h4v14h-4Zm5 0v-5h4v5h-4Z"/>`),
		g.Group(children),
	)
}