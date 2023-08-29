package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfFillSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8A6 6 0 1 1 2 8a6 6 0 0 1 12 0ZM3 8h10A5 5 0 0 0 3 8Z"/>`),
		g.Group(children),
	)
}