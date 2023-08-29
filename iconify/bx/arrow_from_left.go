package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFromLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6h2v12H4zm4 7h8.586l-4.293 4.293l1.414 1.414L20.414 12l-6.707-6.707l-1.414 1.414L16.586 11H8z"/>`),
		g.Group(children),
	)
}