package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15h-2V7H13V5h15a2.002 2.002 0 0 1 2 2Z"/><path fill="currentColor" d="M25 20h-2v-8H8v-2h15a2.002 2.002 0 0 1 2 2Z"/><path fill="currentColor" d="M18 27H4a2.002 2.002 0 0 1-2-2v-8a2.002 2.002 0 0 1 2-2h14a2.002 2.002 0 0 1 2 2v8a2.002 2.002 0 0 1-2 2ZM4 17v8h14.001L18 17Z"/>`),
		g.Group(children),
	)
}