package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M5 0L1 4l4 4l1.5-1.5L4 4l2.5-2.5L5 0z"/>`),
		g.Group(children),
	)
}