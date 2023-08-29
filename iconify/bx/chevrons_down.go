package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 15.586l-4.293-4.293l-1.414 1.414L12 18.414l5.707-5.707l-1.414-1.414z"/><path fill="currentColor" d="m17.707 7.707l-1.414-1.414L12 10.586L7.707 6.293L6.293 7.707L12 13.414z"/>`),
		g.Group(children),
	)
}