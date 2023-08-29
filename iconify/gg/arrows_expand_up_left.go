package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 11H3V3h8v2H6.414l5.364 5.364a1 1 0 0 1-1.414 1.414L5 6.414V11Z"/><path fill-rule="evenodd" d="M19 13a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h4Zm0 2v4h-4v-4h4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}