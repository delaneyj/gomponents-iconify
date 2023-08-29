package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M5 1L2 7h1l3-6H5zM1 2L0 4l1 2h1L1 4l1-2H1zm5 0l1 2l-1 2h1l1-2l-1-2H6z"/>`),
		g.Group(children),
	)
}