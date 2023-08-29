package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventListOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 21v-8h8v8h-8Zm2-2h4v-4h-4v4ZM2 18v-2h9v2H2Zm12-7V3h8v8h-8Zm2-2h4V5h-4v4ZM2 8V6h9v2H2Zm16 9Zm0-10Z"/>`),
		g.Group(children),
	)
}