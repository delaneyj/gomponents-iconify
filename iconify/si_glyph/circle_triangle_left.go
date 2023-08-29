package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTriangleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 0a8 8 0 0 0-8 8a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8zm2 11.846c-.186.205-.774.205-.96 0l-4.9-3.467a.563.563 0 0 1 0-.745l4.9-3.48c.185-.205.773-.205.96 0v7.692z"/>`),
		g.Group(children),
	)
}