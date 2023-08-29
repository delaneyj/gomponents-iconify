package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4.31h-9.37a5 5 0 0 0-5 5v7l-2.92-2.96a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l4.63 4.63a1 1 0 0 0 .33.22a.94.94 0 0 0 .76 0a1.19 1.19 0 0 0 .33-.22L13 14.76a1 1 0 0 0-1.41-1.41l-2.96 2.92v-7a3 3 0 0 1 3-3H21a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}