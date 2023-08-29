package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifebuoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 12a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m12 3l3.35 3.35M9 15l-3.35 3.35m0-12.7L9 9m9.35-3.35L15 9"/></g>`),
		g.Group(children),
	)
}