package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.8 22.7L19.1 21H2v-2h15.1l-4.6-4.6V17h-2v-4.6L6.1 8H2V6l1.4-.7L1.1 3l1.3-1.3l19.7 19.7l-1.3 1.3M4.5 10v7h2v-7h-2m7-6.7L16.7 6H9.2l2 2H21V6l-9.5-5l-4.8 2.5L8.2 5l3.3-1.7m7 12V10h-2v3.3l2 2Z"/>`),
		g.Group(children),
	)
}