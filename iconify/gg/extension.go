package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extension(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 3h8v8h-8V3Zm2 2h4v4h-4V5Z"/><path d="M17 21v-8h-6V7H3v14h14ZM9 9H5v4h4V9ZM5 19v-4h4v4H5Zm6 0v-4h4v4h-4Z"/></g>`),
		g.Group(children),
	)
}