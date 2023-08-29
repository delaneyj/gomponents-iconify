package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertAfterO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 8a1 1 0 0 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V6a1 1 0 1 0-2 0v2H9Z"/><path fill-rule="evenodd" d="M4 9a8 8 0 1 1 16 0A8 8 0 0 1 4 9Zm8 6a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z" clip-rule="evenodd"/><path d="M5 20a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H5Z"/></g>`),
		g.Group(children),
	)
}