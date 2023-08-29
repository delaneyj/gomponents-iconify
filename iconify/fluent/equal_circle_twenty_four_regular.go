package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualCircleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.258 10.5a.75.75 0 0 0 0-1.5H7.77a.75.75 0 0 0 0 1.5h8.488Zm0 4.5a.75.75 0 0 0 0-1.5H7.77a.75.75 0 0 0 0 1.5h8.488ZM22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10s10-4.477 10-10ZM12 3.5a8.5 8.5 0 1 1 0 17a8.5 8.5 0 0 1 0-17Z"/>`),
		g.Group(children),
	)
}