package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfFillSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8A6 6 0 1 1 2 8a6 6 0 0 1 12 0ZM3.5 8h9a4.5 4.5 0 1 0-9 0Z"/>`),
		g.Group(children),
	)
}