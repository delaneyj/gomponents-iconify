package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountLogin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3 0v1h4v5H3v1h5V0H3zm1 2v1H0v1h4v1l2-1.5L4 2z"/>`),
		g.Group(children),
	)
}