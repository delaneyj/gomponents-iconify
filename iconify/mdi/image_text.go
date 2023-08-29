package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13h-8v-2h8v2m0-6h-8v2h8V7m-8 10h8v-2h-8v2m-2-8v6c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V9c0-1.1.9-2 2-2h6c1.1 0 2 .9 2 2m-1.5 6l-2.2-3l-1.8 2.3l-1.2-1.5L3.5 15h7Z"/>`),
		g.Group(children),
	)
}