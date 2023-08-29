package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RomanNumeralFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 7l2 10h2l2-10h-2l-1 5l-1-5h-2m-1 0v2h-1v6h1v2H7v-2h1V9H7V7h4Z"/>`),
		g.Group(children),
	)
}