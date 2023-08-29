package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorPointMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9v6h6V9H9m2 2h2v2h-2v-2m4 7v2h8v-2h-8Z"/>`),
		g.Group(children),
	)
}