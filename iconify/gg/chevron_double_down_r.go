package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleDownR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7.757 8.464L9.172 7.05L12 9.88l2.828-2.829l1.415 1.415L12 12.707L7.757 8.464Z"/><path d="m9.172 11.293l-1.415 1.414L12 16.95l4.243-4.243l-1.415-1.414L12 14.12l-2.828-2.828Z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4V5Zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}