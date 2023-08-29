package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 1L0 5l1.5 1.5L4 4l2.5 2.5L8 5L4 1z"/>`),
		g.Group(children),
	)
}