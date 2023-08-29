package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedEmailOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17V3h18v14H5Zm9-4.725l-7-4.85V15h14V7.425l-7 4.85Zm0-2.425L21 5H7l7 4.85ZM1 21V6.5h2V19h16.5v2H1ZM21 5H7h14Z"/>`),
		g.Group(children),
	)
}