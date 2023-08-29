package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyhole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 12.83a3.001 3.001 0 1 0-2 0V16a1 1 0 1 0 2 0v-3.17ZM11 10a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2ZM4 12a8 8 0 1 0 16 0a8 8 0 0 0-16 0Z"/></g>`),
		g.Group(children),
	)
}