package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyNgn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 9h2V3h2l3.42 6H16V3h2v6h2v2h-2v2h2v2h-2v6h-2l-3.43-6H8v6H6v-6H4v-2h2v-2H4V9m4 0h1.13L8 7.03V9m0 2v2h3.42l-1.14-2H8m8 6v-2h-1.15L16 17m-3.44-6l1.15 2H16v-2h-3.44Z"/>`),
		g.Group(children),
	)
}