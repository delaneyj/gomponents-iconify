package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m14.829 14.828l1.414-1.414L12 9.172l-4.243 4.242l1.415 1.415L12 12l2.829 2.828Z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12Zm11-9a9 9 0 1 0 0 18a9 9 0 0 0 0-18Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}