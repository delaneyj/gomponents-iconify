package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v8h8V0H0zm1 1h6v3L6 3L5 4l2 2v1H6L2 3L1 4V1z"/>`),
		g.Group(children),
	)
}