package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventListSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 21v-8h8v8h-8ZM2 18v-2h9v2H2Zm12-7V3h8v8h-8ZM2 8V6h9v2H2Z"/>`),
		g.Group(children),
	)
}