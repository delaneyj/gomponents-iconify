package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignDirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 12H3.5L6 9.5L3.5 7H11V3l1-1l1 1v4h5l2.5 2.5L18 12h-5v8a2 2 0 0 1 2 2H9a2 2 0 0 1 2-2v-8Z"/>`),
		g.Group(children),
	)
}