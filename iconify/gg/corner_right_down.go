package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.399 6.85a2 2 0 0 0-2.002-1.998l-8 .007l-.002-2l8-.007a4 4 0 0 1 4.004 3.996l.01 10.306l3.78-3.788l1.416 1.412l-6.358 6.37l-6.37-6.358l1.413-1.415l4.119 4.11l-.01-10.635Z"/>`),
		g.Group(children),
	)
}