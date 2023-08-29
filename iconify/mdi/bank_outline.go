package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 10h-2v7h2v-7m6 0h-2v7h2v-7m8.5 9H2v2h19v-2m-2.5-9h-2v7h2v-7m-7-6.74L16.71 6H6.29l5.21-2.74m0-2.26L2 6v2h19V6l-9.5-5Z"/>`),
		g.Group(children),
	)
}