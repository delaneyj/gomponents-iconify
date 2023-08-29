package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalCafeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2h16v2H4Zm0-4V3h18v7h-4v7H4Zm2-2h10V5H6v10Zm12-7h2V5h-2v3ZM6 15h10H6Z"/>`),
		g.Group(children),
	)
}