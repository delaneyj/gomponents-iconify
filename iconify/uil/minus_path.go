package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusPath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.18 4h2.1a1 1 0 0 0 0-2h-2.1a1 1 0 0 0 0 2ZM3 11.28a1 1 0 0 0 1-1v-2.1a1 1 0 0 0-2 0v2.1a1 1 0 0 0 1 1ZM14.46 4a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1h-1a1 1 0 0 0 0 2ZM21 7.54h-4.54a1 1 0 1 0-2 0H8.54a1 1 0 0 0-1 1v5.92a1 1 0 1 0 0 2V21a1 1 0 0 0 1 1H21a1 1 0 0 0 1-1V8.54a1 1 0 0 0-1-1ZM20 20H9.54V9.54H20ZM4 2H3a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0a1 1 0 0 0 0-2Zm0 12.46a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}