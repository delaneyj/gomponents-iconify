package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 1v2l3 3l1.5-1.5L5 4L3 2L2 1H0zm3.41 0l3 3l-1.19 1.22L6 6l2-2l-3-3H3.41zM1.5 2c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5s.22-.5.5-.5z"/>`),
		g.Group(children),
	)
}