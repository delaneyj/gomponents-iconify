package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleLeftO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m12.707 9.172l-1.414-1.415L7.05 12l4.243 4.243l1.414-1.415L9.88 12l2.828-2.828Z"/><path d="m15.536 7.757l1.414 1.415L14.12 12l2.829 2.828l-1.414 1.415L11.293 12l4.243-4.243Z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}