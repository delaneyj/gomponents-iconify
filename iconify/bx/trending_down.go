package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 9.586l-4 4l-6.293-6.293l-1.414 1.414L10 16.414l4-4l4.293 4.293L16 19h6v-6l-2.293 2.293z"/>`),
		g.Group(children),
	)
}