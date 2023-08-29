package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleLeftR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m12.707 9.172l-1.414-1.415L7.05 12l4.243 4.243l1.414-1.415L9.88 12l2.828-2.828Z"/><path d="m15.536 7.757l1.414 1.415L14.12 12l2.829 2.828l-1.414 1.415L11.293 12l4.243-4.243Z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4V5Zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}