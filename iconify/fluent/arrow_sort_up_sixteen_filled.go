package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSortUpSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.22 2.22a.75.75 0 0 1 1.06 0l2.75 2.75a.75.75 0 1 1-1.06 1.06L8.5 4.56v8.69a.75.75 0 0 1-1.5 0V4.56L5.53 6.03a.75.75 0 0 1-1.06-1.06l2.75-2.75Z"/>`),
		g.Group(children),
	)
}