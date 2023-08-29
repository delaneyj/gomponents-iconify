package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDashboardVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5v14h20V5H2m18 7h-4V7h4v5m-6-2h-4V7h4v3m-4 2h4v5h-4v-5M4 7h4v10H4V7m12 10v-3h4v3h-4Z"/>`),
		g.Group(children),
	)
}