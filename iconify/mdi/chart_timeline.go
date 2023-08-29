package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartTimeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h2v18h18v2H2V2m5 8h10v3H7v-3m4 5h10v3H11v-3M6 4h16v4h-2V6H8v2H6V4Z"/>`),
		g.Group(children),
	)
}