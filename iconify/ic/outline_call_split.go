package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineCallSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 4l2.29 2.29l-2.88 2.88l1.42 1.42l2.88-2.88L20 10V4h-6zm-4 0H4v6l2.29-2.29l4.71 4.7V20h2v-8.41l-5.29-5.3L10 4z"/>`),
		g.Group(children),
	)
}