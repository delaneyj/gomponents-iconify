package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUpO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m14.828 12.707l1.415-1.414L12 7.05l-4.243 4.243l1.415 1.414L12 9.88l2.828 2.828Z"/><path d="m14.828 16.95l1.415-1.414L12 11.293l-4.243 4.243l1.415 1.414L12 14.12l2.828 2.829Z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12Zm2 0a9 9 0 1 0 18 0a9 9 0 0 0-18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}