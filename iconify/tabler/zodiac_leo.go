package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZodiacLeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 17a4 4 0 1 0 8 0M3 16a3 3 0 1 0 6 0a3 3 0 1 0-6 0m4-9a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/><path d="M7 7c0 3 2 5 2 9m6-9c0 4-2 6-2 10"/></g>`),
		g.Group(children),
	)
}