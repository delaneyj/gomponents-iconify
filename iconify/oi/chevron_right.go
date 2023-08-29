package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2.5 0L1 1.5L3.5 4L1 6.5L2.5 8l4-4l-4-4z"/>`),
		g.Group(children),
	)
}