package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14.125 13h-4v2h4v-2Z"/><path fill-rule="evenodd" d="M8.125 13a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-1.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm10-.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-1.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.749 14.666A6 6 0 0 0 8.125 18h8c2.44 0 4.54-1.456 5.478-3.547A2.997 2.997 0 0 0 22.875 12c0-1.013-.503-1.91-1.272-2.452A6.001 6.001 0 0 0 16.125 6h-8A6 6 0 0 0 2.75 9.334a3 3 0 0 0 0 5.332ZM8.125 8h8c1.384 0 2.603.702 3.322 1.77c.276.69.428 1.442.428 2.23s-.152 1.54-.428 2.23A3.996 3.996 0 0 1 16.125 16h-8a4 4 0 0 1 0-8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}