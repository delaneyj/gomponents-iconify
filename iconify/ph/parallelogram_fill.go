package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParallelogramFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m246.58 62.57l-64.8 144a16 16 0 0 1-14.59 9.43H24a16 16 0 0 1-14.58-22.57l64.8-144A16 16 0 0 1 88.81 40H232a16 16 0 0 1 14.59 22.57Z"/>`),
		g.Group(children),
	)
}