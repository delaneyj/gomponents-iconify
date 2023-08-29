package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowResetThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.707 3.293a1 1 0 0 1 0 1.414L5.414 8H17.5C23.299 8 28 12.701 28 18.5S23.299 29 17.5 29S7 24.299 7 18.5a1 1 0 1 1 2 0a8.5 8.5 0 1 0 8.5-8.5H5.414l3.293 3.293a1 1 0 1 1-1.414 1.414l-5-5a1 1 0 0 1 0-1.414l5-5a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}