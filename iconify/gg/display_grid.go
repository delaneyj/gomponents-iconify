package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 7v4h4V7H7Zm6 0h4v4h-4V7Zm0 6v4h4v-4h-4Zm-6 0h4v4H7v-4Z"/><path fill-rule="evenodd" d="M3 3h18v18H3V3Zm2 2v14h14V5H5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}