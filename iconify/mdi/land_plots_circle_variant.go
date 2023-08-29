package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandPlotsCircleVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m0 7c0 2.8-2.2 5-5 5s-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5M4 4h4v10H4V4m0 16v-4h4v4H4m16 0H10v-4h10v4Z"/>`),
		g.Group(children),
	)
}