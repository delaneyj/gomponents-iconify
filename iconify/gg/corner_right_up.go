package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.29 10.625L7.877 9.21l6.37-6.358l6.358 6.37l-1.416 1.413l-3.78-3.789l-.01 10.306a4 4 0 0 1-4.004 3.996l-8-.007l.002-2l8 .007a2 2 0 0 0 2.002-1.998l.01-10.636l-4.119 4.111Z"/>`),
		g.Group(children),
	)
}