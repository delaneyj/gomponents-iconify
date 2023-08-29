package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDashboardVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5v14h6V5H2m7 0v5h6V5H9m7 0v9h6V5h-6m-7 6v8h6v-8H9m7 4v4h6v-4h-6Z"/>`),
		g.Group(children),
	)
}