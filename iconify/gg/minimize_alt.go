package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimizeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.073 2l1.415 1.414l-5.85 5.85h2.544v2h-6v-6h2v2.627L20.073 2Zm-8.891 10.264v6h-2v-2.422L3.414 21.61L2 20.196l5.932-5.932h-2.75v-2h6Z"/>`),
		g.Group(children),
	)
}