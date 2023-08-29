package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTriangleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 16a8 8 0 0 0 8-8a8 8 0 0 0-8-8a8 8 0 0 0-8 8a8 8 0 0 0 8 8zM7 4.154c.186-.205.775-.205.96 0l4.9 3.467a.567.567 0 0 1 0 .745l-4.9 3.48c-.185.205-.774.205-.96 0V4.154z"/>`),
		g.Group(children),
	)
}