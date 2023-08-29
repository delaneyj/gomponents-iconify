package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm11 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0ZM6 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm11 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0ZM6 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm11 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/><path fill-rule="evenodd" d="M4 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H4Zm16 2H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}