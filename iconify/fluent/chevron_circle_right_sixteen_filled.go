package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12Zm-.646-8.854l2.5 2.5a.5.5 0 0 1 0 .708l-2.5 2.5a.5.5 0 0 1-.708-.708L8.793 8L6.646 5.854a.5.5 0 1 1 .708-.708Z"/>`),
		g.Group(children),
	)
}