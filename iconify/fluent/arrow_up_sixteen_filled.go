package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 14a.75.75 0 0 1-.75-.75V4.463L4.309 7.75a.75.75 0 0 1-1.118-1L7.441 2A.75.75 0 0 1 8.56 2l4.25 4.75a.75.75 0 1 1-1.118 1L8.75 4.463v8.787A.75.75 0 0 1 8 14Z"/>`),
		g.Group(children),
	)
}