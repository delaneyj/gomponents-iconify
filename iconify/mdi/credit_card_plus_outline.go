package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18h3v2h-3v3h-2v-3h-3v-2h3v-3h2v3M19 8V6H3v2h16m0 4H3v6h11v2H3a2 2 0 0 1-2-2V6c0-1.11.89-2 2-2h16a2 2 0 0 1 2 2v7h-2v-1Z"/>`),
		g.Group(children),
	)
}