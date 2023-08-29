package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.71 10.625l1.413-1.415l-6.37-6.358l-6.358 6.37l1.416 1.413l3.78-3.789l.01 10.306a4 4 0 0 0 4.004 3.996l8-.007l-.002-2l-8 .007a2 2 0 0 1-2.002-1.998l-.01-10.636l4.119 4.111Z"/>`),
		g.Group(children),
	)
}