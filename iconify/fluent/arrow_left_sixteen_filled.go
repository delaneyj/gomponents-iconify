package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8a.75.75 0 0 1-.75.75H4.463l3.287 2.941a.75.75 0 1 1-1 1.118L2 8.559A.75.75 0 0 1 2 7.44l4.75-4.25a.75.75 0 1 1 1 1.118L4.463 7.25h8.787A.75.75 0 0 1 14 8Z"/>`),
		g.Group(children),
	)
}