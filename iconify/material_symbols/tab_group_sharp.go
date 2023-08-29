package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabGroupSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V6h2v14h14v2H2Zm4-4V2h16v16H6Zm7-10h7V4h-7v4Z"/>`),
		g.Group(children),
	)
}