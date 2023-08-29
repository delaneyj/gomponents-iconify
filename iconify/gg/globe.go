package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.853 8a5 5 0 1 1 10 0a5 5 0 0 1-10 0Zm5 3a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z" clip-rule="evenodd"/><path d="M5 12.13a8.001 8.001 0 0 0 5.941 3.819V18H8.853v2h6v-2h-1.912v-2.073A8.002 8.002 0 0 0 18.63 3.745l-1.704 1.048a6 6 0 1 1-10.221 6.288L5 12.13Z"/></g>`),
		g.Group(children),
	)
}