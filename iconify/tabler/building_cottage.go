package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingCottage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 21h18M4 21V10l2.5-4.5L12 3l5.5 2.5L20 10v11"/><path d="M10 9a2 2 0 1 0 4 0a2 2 0 1 0-4 0M9 21v-5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5"/></g>`),
		g.Group(children),
	)
}