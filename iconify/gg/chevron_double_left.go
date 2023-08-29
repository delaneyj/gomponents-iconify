package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M18.364 7.757L16.95 6.343L11.293 12l5.657 5.657l1.414-1.414L14.12 12l4.243-4.243Z"/><path d="m11.293 6.343l1.414 1.414L8.464 12l4.243 4.243l-1.414 1.414L5.636 12l5.657-5.657Z"/></g>`),
		g.Group(children),
	)
}