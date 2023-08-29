package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm5-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/><path fill-rule="evenodd" d="M4 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H4Zm16 2H4a1 1 0 0 0-1 1v1h18V6a1 1 0 0 0-1-1ZM3 18V9h18v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}