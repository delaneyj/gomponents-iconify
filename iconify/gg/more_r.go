package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm7-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm3 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path fill-rule="evenodd" d="M0 5a3 3 0 0 1 3-3h18a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3V5Zm3-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}