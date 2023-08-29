package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Debug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 11a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1Zm1 3a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2Z"/><path fill-rule="evenodd" d="M9.094 4.75A3.986 3.986 0 0 1 8 2h2a2 2 0 1 0 4 0h2a3.986 3.986 0 0 1-1.095 2.75A6.02 6.02 0 0 1 17.66 8H19a1 1 0 1 1 0 2h-1v2h1a1 1 0 1 1 0 2h-1v2h1a1 1 0 1 1 0 2h-1.341A6.003 6.003 0 0 1 6.34 18H5a1 1 0 1 1 0-2h1v-2H5a1 1 0 1 1 0-2h1v-2H5a1 1 0 1 1 0-2h1.341a6.02 6.02 0 0 1 2.753-3.25ZM8 16v-6a4 4 0 1 1 8 0v6a4 4 0 0 1-8 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}