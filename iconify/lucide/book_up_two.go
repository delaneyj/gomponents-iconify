package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2"/><path d="M18 2h2v20H6.5a2.5 2.5 0 0 1 0-5H20m-8-4V7"/><path d="m9 10l3-3l3 3M9 5l3-3l3 3"/></g>`),
		g.Group(children),
	)
}