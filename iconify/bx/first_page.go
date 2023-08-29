package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.293 17.707l1.414-1.414L13.414 12l4.293-4.293l-1.414-1.414L10.586 12zM7 6h2v12H7z"/>`),
		g.Group(children),
	)
}