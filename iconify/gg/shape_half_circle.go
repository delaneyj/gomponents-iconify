package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeHalfCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4a8 8 0 1 0 0 16v-3a5 5 0 0 1 0-10V4Z"/>`),
		g.Group(children),
	)
}