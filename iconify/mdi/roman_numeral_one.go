package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RomanNumeralOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 7v2h-1v6h1v2h-4v-2h1V9h-1V7h4Z"/>`),
		g.Group(children),
	)
}