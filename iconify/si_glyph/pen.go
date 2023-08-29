package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.224-.387l1.869 1.868l-3.035 3.035l-1.87-1.869zM11.252 3.48l1.876 1.876l-7.115 7.115l-1.876-1.876zM1.021 15.957l2.143-4.109l1.877 1.877l-4.02 2.232zm10.534-5.996l-.468-.477l2.965-2.919l-.393-.434l.496-.448l.822.91l-3.422 3.368z"/>`),
		g.Group(children),
	)
}