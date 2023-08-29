package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowToRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6h2v12h-2zM4 13h8.586l-4.293 4.293l1.414 1.414L16.414 12L9.707 5.293L8.293 6.707L12.586 11H4z"/>`),
		g.Group(children),
	)
}