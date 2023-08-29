package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.293 11.293l1.414 1.414L12 8.414l4.293 4.293l1.414-1.414L12 5.586z"/><path fill="currentColor" d="m6.293 16.293l1.414 1.414L12 13.414l4.293 4.293l1.414-1.414L12 10.586z"/>`),
		g.Group(children),
	)
}