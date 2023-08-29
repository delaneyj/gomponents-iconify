package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12Zm.646-8.854l-2.5 2.5a.5.5 0 0 0 0 .708l2.5 2.5a.5.5 0 0 0 .708-.708L7.207 8l2.147-2.146a.5.5 0 1 0-.708-.708Z"/>`),
		g.Group(children),
	)
}