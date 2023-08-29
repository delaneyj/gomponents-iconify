package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MortarPestlePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 14h-3v3h-2v-3H8v-2h3V9h2v3h3m5-7h-2.65l1.15-3.15L17.15 1l-1.46 4H3v2l2 6l-2 6v2h18v-2l-2-6l2-6V5Z"/>`),
		g.Group(children),
	)
}