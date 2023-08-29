package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1.5 1L0 2.5l4 4l4-4L6.5 1L4 3.5L1.5 1z"/>`),
		g.Group(children),
	)
}