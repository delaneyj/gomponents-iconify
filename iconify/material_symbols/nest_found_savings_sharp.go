package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestFoundSavingsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l-3-3H3V2h18v18h-6l-3 3Zm0-7q2.15 0 3.575-1.5T17 11V6h-5Q9.975 6 8.487 7.425T7 11q0 .75.213 1.425t.587 1.25l-1.1 1.1l1.35 1.35l1.025-1q.675.425 1.4.65T12 16Zm-2.75-3.775l3.3-3.3l1.425 1.4l-3.4 3.325l-1.325-1.425Z"/>`),
		g.Group(children),
	)
}