package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFromRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6h2v12h-2zm-2 5H7.414l4.293-4.293l-1.414-1.414L3.586 12l6.707 6.707l1.414-1.414L7.414 13H16z"/>`),
		g.Group(children),
	)
}