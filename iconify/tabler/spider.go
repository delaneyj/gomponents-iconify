package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 4v2l5 5M2.5 9.5L4 11h6m-6 8v-2l6-6m9-7v2l-5 5m7.5-1.5L20 11h-6m6 8v-2l-6-6"/><path d="M8 15a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/><path d="M10 9a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/></g>`),
		g.Group(children),
	)
}