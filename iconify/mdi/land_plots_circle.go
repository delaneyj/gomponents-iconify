package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandPlotsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-7 5.5C13 5.6 14.6 4 16.5 4S20 5.6 20 7.5S18.4 11 16.5 11S13 9.4 13 7.5M4 4h7v7H4V4m0 16v-7h7v7H4m16 0h-7v-7h7v7Z"/>`),
		g.Group(children),
	)
}