package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RomanNumeralSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 7l2 10h2l2-10h-2l-1 5l-1-5H6m11 0v2h-1v6h1v2h-4v-2h1V9h-1V7h4Z"/>`),
		g.Group(children),
	)
}