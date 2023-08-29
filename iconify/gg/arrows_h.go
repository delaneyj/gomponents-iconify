package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.657 9.172L4.243 7.757L0 12l4.243 4.243l1.414-1.415L3.829 13H10v-2H3.83l1.828-1.828ZM14 11v2h6.172l-1.829 1.828l1.414 1.415L24 12l-4.243-4.243l-1.414 1.415L20.172 11H14Z"/>`),
		g.Group(children),
	)
}