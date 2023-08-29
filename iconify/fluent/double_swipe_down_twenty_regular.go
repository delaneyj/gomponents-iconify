package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleSwipeDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.854 17.854a.5.5 0 0 1-.708 0l-3-3a.5.5 0 0 1 .708-.707L5 16.293V5.5a.5.5 0 0 1 1 0v10.793l2.146-2.146a.5.5 0 0 1 .708.707l-3 3Zm9 0a.5.5 0 0 1-.708 0l-3-3a.5.5 0 0 1 .708-.707L14 16.293V5.5a.5.5 0 0 1 1 0v10.793l2.146-2.146a.5.5 0 0 1 .708.707l-3 3ZM11 5.5a3.5 3.5 0 0 0 2 3.163V7.5a2.5 2.5 0 1 1 3 0v1.163A3.5 3.5 0 1 0 11 5.5ZM4 8.663a3.5 3.5 0 1 1 3 0V7.5a2.5 2.5 0 1 0-3 0v1.163Z"/>`),
		g.Group(children),
	)
}