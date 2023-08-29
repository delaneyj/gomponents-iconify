package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.457 23.207a1 1 0 0 1-1.414 0l-8.75-8.75a1 1 0 0 1 0-1.414l8.75-8.75a1 1 0 1 1 1.414 1.414l-8.043 8.043l8.043 8.043a1 1 0 0 1 0 1.414Z"/>`),
		g.Group(children),
	)
}