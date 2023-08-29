package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 5h-3V3a1 1 0 0 0-2 0v2H6a1 1 0 0 0-.78.38l-2 2.5a1 1 0 0 0 0 1.24l2 2.5A1 1 0 0 0 6 12h5v8H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2v-8h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Zm-1 5H6.48l-1.2-1.5L6.48 7H15Z"/>`),
		g.Group(children),
	)
}