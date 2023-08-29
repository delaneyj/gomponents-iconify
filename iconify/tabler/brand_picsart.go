package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandPicsart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 9a7 7 0 1 0 14 0A7 7 0 1 0 5 9"/><path d="M9 9a3 3 0 1 0 6 0a3 3 0 1 0-6 0M5 9v11a2 2 0 1 0 4 0v-4.5"/></g>`),
		g.Group(children),
	)
}