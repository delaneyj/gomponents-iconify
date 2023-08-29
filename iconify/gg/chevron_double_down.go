package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7.757 5.636L6.343 7.05L12 12.707l5.657-5.657l-1.414-1.414L12 9.88L7.757 5.636Z"/><path d="m6.343 12.707l1.414-1.414L12 15.536l4.243-4.243l1.414 1.414L12 18.364l-5.657-5.657Z"/></g>`),
		g.Group(children),
	)
}