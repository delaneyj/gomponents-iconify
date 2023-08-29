package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingCarousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 12a6 6 0 1 0 12 0a6 6 0 1 0-12 0"/><path d="M3 8a2 2 0 1 0 4 0a2 2 0 1 0-4 0m7-4a2 2 0 1 0 4 0a2 2 0 1 0-4 0m7 4a2 2 0 1 0 4 0a2 2 0 1 0-4 0M3 16a2 2 0 1 0 4 0a2 2 0 1 0-4 0m14 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-9 6l4-10l4 10"/></g>`),
		g.Group(children),
	)
}