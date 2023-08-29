package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.293 13.293l1.414 1.414L12 10.414l4.293 4.293l1.414-1.414L12 7.586z"/>`),
		g.Group(children),
	)
}