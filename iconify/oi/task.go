package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Task(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v7h7V3.41l-1 1V6H1V1h3.59l1-1H0zm7 0L4 3L3 2L2 3l2 2l4-4l-1-1z"/>`),
		g.Group(children),
	)
}