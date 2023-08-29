package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v2h8V0H0zm2 3l2 2l2-2H2zM0 7v1h8V7H0z"/>`),
		g.Group(children),
	)
}