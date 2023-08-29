package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleDownO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7.757 8.464L9.172 7.05L12 9.88l2.828-2.829l1.415 1.415L12 12.707L7.757 8.464Z"/><path d="m9.172 11.293l-1.415 1.414L12 16.95l4.243-4.243l-1.415-1.414L12 14.12l-2.828-2.828Z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}