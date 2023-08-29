package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowULeftTopBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 21H6v-4h7.5c1.93 0 3.5-1.57 3.5-3.5S15.43 10 13.5 10H11v4L4 8l7-6v4h2.5c4.14 0 7.5 3.36 7.5 7.5S17.64 21 13.5 21Z"/>`),
		g.Group(children),
	)
}