package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleRightO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.464 7.757L7.05 9.172L9.88 12l-2.83 2.828l1.415 1.415L12.707 12L8.464 7.757Z"/><path d="m11.293 9.172l1.414-1.415L16.95 12l-4.243 4.243l-1.414-1.415L14.12 12l-2.828-2.828Z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12Zm2 0a9 9 0 1 0 18 0a9 9 0 0 0-18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}