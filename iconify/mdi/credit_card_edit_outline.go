package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardEditOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 18.9l6.1-6.1l2.1 2.1l-6.1 6.1H13v-2.1m8.4-7.6l1.3 1.3c.2.2.2.5 0 .7l-1 1l-2.1-2l1-1c.1-.1.2-.2.4-.2s.3.1.4.2M11 18H4v-6h13.1L22 7.1V6c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h7v-2M4 6h16v2H4V6Z"/>`),
		g.Group(children),
	)
}