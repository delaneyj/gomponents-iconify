package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowToLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6h2v12H4zm10.293-.707L7.586 12l6.707 6.707l1.414-1.414L11.414 13H20v-2h-8.586l4.293-4.293z"/>`),
		g.Group(children),
	)
}