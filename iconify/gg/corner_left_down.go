package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.601 6.85a2 2 0 0 1 2.002-1.998l8 .007l.002-2l-8-.007a4 4 0 0 0-4.004 3.996l-.01 10.306l-3.78-3.788l-1.416 1.412l6.358 6.37l6.37-6.358l-1.413-1.415l-4.119 4.11l.01-10.635Z"/>`),
		g.Group(children),
	)
}