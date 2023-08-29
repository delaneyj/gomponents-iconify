package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m14.854 11.974l1.415-1.414l-4.243-4.243l-4.243 4.243l1.414 1.414l1.829-1.828v7.537h2v-7.537l1.828 1.828Z"/><path fill-rule="evenodd" d="M1 19a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14Zm4 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}