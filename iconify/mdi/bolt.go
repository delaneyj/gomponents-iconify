package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17.7V21h-4v-.7l4-2.6M17 3H7v3h10V3m-2 4l-1 .7V7h-4v3.3L9 11v1l6-3.9V7m0 4l-1 .7v-2l-4 2.7v2L9 15v1l6-3.9V11m0 4l-1 .7v-2l-4 2.7v2L9 19v1l6-3.9V15Z"/>`),
		g.Group(children),
	)
}