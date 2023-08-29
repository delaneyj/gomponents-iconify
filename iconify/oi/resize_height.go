package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.5 0L1 3h2v2H1l2.5 3L6 5H4V3h2L3.5 0z"/>`),
		g.Group(children),
	)
}