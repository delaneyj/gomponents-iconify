package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m17.657 11.293l-1.414 1.414L12 8.464l-4.243 4.243l-1.414-1.414L12 5.636l5.657 5.657Z"/><path d="m17.657 16.95l-1.414 1.414L12 14.12l-4.243 4.243l-1.414-1.414L12 11.293l5.657 5.657Z"/></g>`),
		g.Group(children),
	)
}