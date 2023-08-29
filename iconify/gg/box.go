package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 12a1 1 0 1 0 0 2h4a1 1 0 0 0 0-2h-4Z"/><path fill-rule="evenodd" d="M4 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H4Zm16 2H4a1 1 0 0 0-1 1v3h18V5a1 1 0 0 0-1-1ZM3 19v-9h18v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}