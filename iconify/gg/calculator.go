package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 5H7v2h10V5ZM7 9h2v2H7V9Zm2 4H7v2h2v-2Zm-2 4h2v2H7v-2Zm6-8h-2v2h2V9Zm-2 4h2v2h-2v-2Zm2 4h-2v2h2v-2Zm2-8h2v2h-2V9Zm2 4h-2v6h2v-6Z"/><path fill-rule="evenodd" d="M3 3a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3Zm2 0h14v18H5V3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}