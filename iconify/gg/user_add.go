package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/><path d="M11 14a1 1 0 0 1 1 1v6h2v-6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v6h2v-6a1 1 0 0 1 1-1h6Zm7-7h2v2h2v2h-2v2h-2v-2h-2V9h2V7Z"/></g>`),
		g.Group(children),
	)
}