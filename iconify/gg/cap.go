package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8 18v2h8v-2H8Z"/><path fill-rule="evenodd" d="M13.988 3.22a2 2 0 1 0-3.976 0a9.003 9.003 0 0 0-6.94 9.926A3.001 3.001 0 0 0 1 16v4a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-4c0-1.333-.87-2.463-2.072-2.854a9.003 9.003 0 0 0-6.94-9.926ZM12 5a7 7 0 0 0-6.93 8h13.86A7 7 0 0 0 12 5ZM3 16a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}