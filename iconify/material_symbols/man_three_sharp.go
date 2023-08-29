package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManThreeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22.25v-7H8v-8h8v8h-2v7h-4ZM12 6.5L9.75 4.25L12 2l2.25 2.25L12 6.5Z"/>`),
		g.Group(children),
	)
}