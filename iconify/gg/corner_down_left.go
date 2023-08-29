package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.15 13.4a2 2 0 0 0 2-2v-8h2v8a4 4 0 0 1-4 4H6.844l3.785 3.785L9.214 20.6L2.85 14.235l6.364-6.364l1.415 1.415L6.514 13.4H17.15Z"/>`),
		g.Group(children),
	)
}