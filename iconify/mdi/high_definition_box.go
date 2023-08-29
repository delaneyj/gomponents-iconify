package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighDefinitionBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2m-8 12H9.5v-2h-2v2H6V9h1.5v2.5h2V9H11v6m2-6h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4V9m1.5 4.5h2v-3h-2v3Z"/>`),
		g.Group(children),
	)
}