package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExternalLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v8h8V6H7v1H1V1h1V0H0zm4 0l1.5 1.5L3 4l1 1l2.5-2.5L8 4V0H4z"/>`),
		g.Group(children),
	)
}