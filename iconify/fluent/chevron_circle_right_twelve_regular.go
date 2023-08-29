package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleRightTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.854 4.146a.5.5 0 1 0-.708.708L6.293 6L5.146 7.146a.5.5 0 1 0 .708.708l1.5-1.5a.5.5 0 0 0 0-.708l-1.5-1.5ZM6 1a5 5 0 1 0 0 10A5 5 0 0 0 6 1ZM2 6a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}