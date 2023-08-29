package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16v2H6V2h2v11.57l5.71-9l3.16 2.11l2.42-2.42l1.42 1.42l-3.58 3.61l-2.84-1.89L8.82 16M4 20V4H2v18h20v-2Z"/>`),
		g.Group(children),
	)
}