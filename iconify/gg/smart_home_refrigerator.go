package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartHomeRefrigerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 6a1 1 0 0 1 2 0v2a1 1 0 1 1-2 0V6Zm1 7a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1Z"/><path fill-rule="evenodd" d="M5 4a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4Zm3-1h8a1 1 0 0 1 1 1v6H7V4a1 1 0 0 1 1-1Zm-1 9h10v8a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}