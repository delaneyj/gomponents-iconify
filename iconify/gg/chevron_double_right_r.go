package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleRightR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.464 7.757L7.05 9.172L9.88 12l-2.83 2.828l1.415 1.415L12.707 12L8.464 7.757Z"/><path d="m11.293 9.172l1.414-1.415L16.95 12l-4.243 4.243l-1.414-1.415L14.12 12l-2.828-2.828Z"/><path fill-rule="evenodd" d="M23 5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5Zm-4-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}