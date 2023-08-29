package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeGroupMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6H1l4-4l4 4H8v3H6V6H4v3H2V6m11 4.9l1.3 1.1H16V9h2v3h3V8h1l-5-5l-5 5h1v2.9m.8 11.1c-.5-.9-.8-1.9-.8-3c0-1.6.6-3.1 1.7-4.1L9 10l-7 6h2v6h3v-5h4v5h2.8m1.2-4v2h8v-2h-8Z"/>`),
		g.Group(children),
	)
}