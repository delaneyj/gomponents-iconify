package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalDistribute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2v2H2V2h20M7 10.5v3h10v-3H7M2 20v2h20v-2H2Z"/>`),
		g.Group(children),
	)
}