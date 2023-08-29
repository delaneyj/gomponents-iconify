package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 1v1l4 2l4-2V1H0zm0 2v4h8V3L4 5L0 3z"/>`),
		g.Group(children),
	)
}