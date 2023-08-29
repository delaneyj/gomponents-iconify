package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampDesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 5l-3 3l2 7l8-8l-7-2Z"/><path d="m14 5l-3 3l-3-3l3-3l3 3Z"/><path d="M9.5 6.5L4 12l3 6m-4 4v-2c0-1.1.9-2 2-2h4a2 2 0 0 1 2 2v2H3Z"/></g>`),
		g.Group(children),
	)
}