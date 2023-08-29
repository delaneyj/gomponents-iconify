package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 6a1 1 0 1 0-2 0v3H7a1 1 0 0 0 0 2h4v7a1 1 0 1 0 2 0v-7h4a1 1 0 1 0 0-2h-4V6Z"/>`),
		g.Group(children),
	)
}