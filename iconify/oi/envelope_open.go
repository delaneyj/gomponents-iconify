package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0L0 2v6h8V2L4 0zm0 1.13l3 1.5v1.88l-3 1.5l-3-1.5V2.63l3-1.5zM2 3.01v1l2 1l2-1v-1H2z"/>`),
		g.Group(children),
	)
}