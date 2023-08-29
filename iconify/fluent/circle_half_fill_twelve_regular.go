package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfFillTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a4 4 0 1 1 8 0H2Zm4-5a5 5 0 1 0 0 10A5 5 0 0 0 6 1Z"/>`),
		g.Group(children),
	)
}