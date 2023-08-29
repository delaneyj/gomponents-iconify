package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 21H2V3h2v16h2v-2h4v2h2v-3h4v3h2v-2h4v4m-4-7h4v2h-4v-2m-6-8h4v3h-4V6m4 9h-4v-5h4v5M6 10h4v2H6v-2m4 6H6v-3h4v3Z"/>`),
		g.Group(children),
	)
}