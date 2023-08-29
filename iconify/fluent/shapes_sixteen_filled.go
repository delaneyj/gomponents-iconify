package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 5.5A4.5 4.5 0 0 1 9.973 5H8.5A3.5 3.5 0 0 0 5 8.5v1.473A4.5 4.5 0 0 1 1 5.5Zm7.5.5A2.5 2.5 0 0 0 6 8.5v4A2.5 2.5 0 0 0 8.5 15h4a2.5 2.5 0 0 0 2.5-2.5v-4A2.5 2.5 0 0 0 12.5 6h-4Z"/>`),
		g.Group(children),
	)
}