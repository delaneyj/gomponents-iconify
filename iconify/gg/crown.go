package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m2.5 6.091l4.72 4.72L12 6.031l4.781 4.78L21.5 6.092v8.877a3 3 0 0 1-3 3h-13a3 3 0 0 1-3-3V6.091Zm17 4.818v4.06a1 1 0 0 1-1 1h-13a1 1 0 0 1-1-1v-4.061l2.72 2.72L12 8.848l4.781 4.78l2.719-2.72Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}