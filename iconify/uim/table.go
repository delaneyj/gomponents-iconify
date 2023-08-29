package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 22H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1ZM4 20h16V4H4Z"/><path fill="currentColor" d="M9 22a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v18a1 1 0 0 1-1 1zm6 0a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v18a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M21 10H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 6H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}