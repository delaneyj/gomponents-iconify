package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertBeforeO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 3a1 1 0 0 1 0-2h14a1 1 0 1 1 0 2H5Zm4 12a1 1 0 1 1 0-2h2v-2a1 1 0 1 1 2 0v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9Z"/><path fill-rule="evenodd" d="M4 14a8 8 0 1 0 16 0a8 8 0 0 0-16 0Zm8-6a6 6 0 1 0 0 12a6 6 0 0 0 0-12Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}