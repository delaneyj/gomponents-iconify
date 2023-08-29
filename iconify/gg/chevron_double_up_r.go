package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUpR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m14.828 12.481l1.415-1.414L12 6.824l-4.243 4.243l1.415 1.414L12 9.653l2.828 2.828Z"/><path d="m14.828 16.724l1.415-1.414L12 11.067L7.757 15.31l1.415 1.414L12 13.895l2.828 2.829Z"/><path fill-rule="evenodd" d="M23 4.774a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4v-14Zm-4-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-14a2 2 0 0 0-2-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}