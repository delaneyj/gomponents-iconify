package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.5 0L2 3h2l-.66 2H2l1 3l3-3H4.5L6 2H4l1-2H3.5z"/>`),
		g.Group(children),
	)
}