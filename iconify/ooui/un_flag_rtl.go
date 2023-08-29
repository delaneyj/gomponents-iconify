package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnFlagRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m0 1.2l4.27 4.27L0 7l11.84 6.04l.16.16V20h2v-4.8l4.74 4.74l1.198-1.198L1.198.002zM14 2L7.809 4.209L14 10.399z"/>`),
		g.Group(children),
	)
}