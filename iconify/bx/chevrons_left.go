package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.707 7.707l-1.414-1.414L5.586 12l5.707 5.707l1.414-1.414L8.414 12z"/><path fill="currentColor" d="M16.293 6.293L10.586 12l5.707 5.707l1.414-1.414L13.414 12l4.293-4.293z"/>`),
		g.Group(children),
	)
}