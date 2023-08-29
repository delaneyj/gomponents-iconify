package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageFlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v18l4-4h9v-8h7V4m.5 12h-2.2l1.7-4h-5v6h2v5l3.5-7Z"/>`),
		g.Group(children),
	)
}