package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveGroupSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18v-4h2v2h12V6H8v2H6V2h16v16H6Zm-4 4V6h2v14h14v2H2Zm11-7l-1.4-1.4l1.575-1.6H6v-2h7.175L11.6 8.4L13 7l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}