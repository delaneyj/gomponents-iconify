package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v5.91c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09V2H5.02v1.03H2.99V2h-3z"/>`),
		g.Group(children),
	)
}