package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m246 117.35l-36 40a8 8 0 0 1-6 2.65h-68v64a8 8 0 0 1-16 0v-64H40a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h80V32a8 8 0 0 1 16 0v32h68a8 8 0 0 1 6 2.65l36 40a8 8 0 0 1 0 10.7Z"/>`),
		g.Group(children),
	)
}