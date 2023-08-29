package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 8.041a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-16 0zM12.846 6c.205.185.205.772 0 .96l-3.467 4.9a.568.568 0 0 1-.746 0l-3.479-4.9c-.205-.187-.205-.774 0-.96h7.692z"/>`),
		g.Group(children),
	)
}