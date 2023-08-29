package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGrowthAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 20H4v-2h3a1 1 0 0 0 0-2H4v-2h7a1 1 0 0 0 0-2H4v-2h11a1 1 0 0 0 0-2H4V6h15a1 1 0 0 0 0-2H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}