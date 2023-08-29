package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3.75A.75.75 0 0 1 8.75 3h7.5a.75.75 0 0 1 .75.75v7.5a.75.75 0 1 1-1.5 0V5.56L4.28 16.78a.75.75 0 0 1-1.06-1.06L14.44 4.5H8.75A.75.75 0 0 1 8 3.75Z"/>`),
		g.Group(children),
	)
}