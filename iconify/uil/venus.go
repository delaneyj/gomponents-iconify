package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Venus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 9a7 7 0 1 0-8 6.92V18h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1v-2.08A7 7 0 0 0 19 9Zm-7 5a5 5 0 1 1 5-5a5 5 0 0 1-5 5Z"/>`),
		g.Group(children),
	)
}