package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingPavilion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21h7v-3a2 2 0 0 1 4 0v3h7M6 21v-9m12 9v-9M6 12h12a3 3 0 0 0 3-3a9 8 0 0 1-9-6a9 8 0 0 1-9 6a3 3 0 0 0 3 3"/>`),
		g.Group(children),
	)
}