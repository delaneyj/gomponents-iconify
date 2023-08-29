package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/><path d="M11 14a1 1 0 0 1 1 1v6h2v-6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v6h2v-6a1 1 0 0 1 1-1h6Zm11-5h-6v2h6V9Z"/></g>`),
		g.Group(children),
	)
}