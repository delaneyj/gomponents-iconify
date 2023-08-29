package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DismissCircleTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.854 4.146a.5.5 0 0 1 0 .708L6.707 6l1.147 1.146a.5.5 0 1 1-.708.708L6 6.707L4.854 7.854a.5.5 0 1 1-.708-.708L5.293 6L4.146 4.854a.5.5 0 1 1 .708-.708L6 5.293l1.146-1.147a.5.5 0 0 1 .708 0ZM6 1a5 5 0 1 0 0 10A5 5 0 0 0 6 1ZM2 6a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}