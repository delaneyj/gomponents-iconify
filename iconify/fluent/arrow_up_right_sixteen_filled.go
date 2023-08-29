package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 2.75A.75.75 0 0 1 7.75 2h5.5a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-1.5 0V4.56l-9.22 9.22a.75.75 0 0 1-1.06-1.06l9.22-9.22H7.75A.75.75 0 0 1 7 2.75Z"/>`),
		g.Group(children),
	)
}