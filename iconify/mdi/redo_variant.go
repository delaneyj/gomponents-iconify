package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 7A6.5 6.5 0 0 0 4 13.5a6.5 6.5 0 0 0 6.5 6.5H14v-2h-3.5C8 18 6 16 6 13.5S8 9 10.5 9h5.67l-3.08 3.09l1.41 1.41L20 8l-5.5-5.5l-1.42 1.41L16.17 7H10.5M18 18h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}