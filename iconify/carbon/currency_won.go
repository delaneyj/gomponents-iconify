package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyWon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 14v-2h-4.955L24 5h-2l-2 18l-3-15h-2l-3 15l-2-18H8l.955 7H4v2h5.227l.409 3H4v2h5.909L11 27h2l3-15l3 15h2l1.091-8H28v-2h-5.636l.409-3H28z"/>`),
		g.Group(children),
	)
}