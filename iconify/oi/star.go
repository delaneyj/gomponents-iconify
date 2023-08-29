package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0L3 3H0l2.5 2l-1 3L4 6l2.5 2l-1-3L8 3H5L4 0z"/>`),
		g.Group(children),
	)
}