package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventSeatOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-6h16v6h-2v-4H6v4H4Zm-1-7v-3h3v3H3Zm4 0V3h10v11H7Zm11 0v-3h3v3h-3Zm-9-2h6V5H9v7Zm0 0h6h-6Z"/>`),
		g.Group(children),
	)
}