package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14h10l-3.45-4.5l-2.3 3l-1.55-2L9 14Zm-3 4V2h16v16H6Zm-4 4V6h2v14h14v2H2Z"/>`),
		g.Group(children),
	)
}