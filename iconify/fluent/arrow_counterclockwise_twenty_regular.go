package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterclockwiseTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.934 9.05A7 7 0 0 0 4.377 5.83L4.25 6H7.5a.5.5 0 1 1 0 1h-4a.5.5 0 0 1-.5-.5V2.502a.5.5 0 1 1 1 0v2.207a8 8 0 1 1-1.986 4.775a.5.5 0 1 1 .998.064a7 7 0 1 0 13.922-.496Z"/>`),
		g.Group(children),
	)
}