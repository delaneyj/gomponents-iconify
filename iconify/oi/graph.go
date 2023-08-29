package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Graph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M7.03 0L4 3L3 2L0 5.03l1 1L3 4l1 1l4-4l-.97-1zM0 7v1h8V7H0z"/>`),
		g.Group(children),
	)
}