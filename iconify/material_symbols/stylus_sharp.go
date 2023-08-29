package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.45 20.1L2.7 21.3l1.2-5.75l4.55 4.55Zm1.625-1.225l-4.95-4.95l11.75-11.75l4.95 4.95l-11.75 11.75Z"/>`),
		g.Group(children),
	)
}