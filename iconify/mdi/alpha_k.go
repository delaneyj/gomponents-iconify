package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7v10h2v-3.33L13 17h2l-3-5l3-5h-2l-2 3.33V7H9Z"/>`),
		g.Group(children),
	)
}