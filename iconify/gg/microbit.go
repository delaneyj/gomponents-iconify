package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm12-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path fill-rule="evenodd" d="M7 5a7 7 0 0 0 0 14h10a7 7 0 1 0 0-14H7Zm10 3H7a4 4 0 1 0 0 8h10a4 4 0 0 0 0-8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}