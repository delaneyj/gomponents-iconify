package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileUpside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16 11h-2a2 2 0 1 0-4 0H8a4 4 0 1 1 8 0Zm-6 3a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm5-1a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/><path fill-rule="evenodd" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10s10-4.477 10-10Zm-2 0a8 8 0 1 0-16 0a8 8 0 0 0 16 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}