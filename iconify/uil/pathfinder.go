package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pathfinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 14.46a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 0-2ZM8.18 4h2.1a1 1 0 0 0 0-2h-2.1a1 1 0 0 0 0 2Zm6.28 0a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1h-1a1 1 0 0 0 0 2ZM4 2H3a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0a1 1 0 0 0 0-2Zm-1 9.28a1 1 0 0 0 1-1v-2.1a1 1 0 0 0-2 0v2.1a1 1 0 0 0 1 1ZM15.82 20h-2.1a1 1 0 1 0 0 2h2.1a1 1 0 0 0 0-2ZM21 7.54h-1a1 1 0 0 0 0 2a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1Zm0 5.18a1 1 0 0 0-1 1v2.1a1 1 0 0 0 2 0v-2.1a1 1 0 0 0-1-1Zm-4.54-5.18a1 1 0 1 0-2 0H8.54a1 1 0 0 0-1 1v5.92a1 1 0 1 0 0 2a1 1 0 0 0 2 0h5.92a1 1 0 0 0 1-1V9.54a1 1 0 1 0 0-2Zm-2 6.92H9.54V9.54h4.92ZM21 19a1 1 0 0 0-1 1a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1ZM9.54 20a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}