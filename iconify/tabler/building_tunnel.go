package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingTunnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 21h14a2 2 0 0 0 2-2v-7a9 9 0 0 0-18 0v7a2 2 0 0 0 2 2z"/><path d="M8 21v-9a4 4 0 1 1 8 0v9M3 17h4m10 0h4m0-5h-4M7 12H3m9-9v5M6 6l3 3m6 0l3-3l-3 3z"/></g>`),
		g.Group(children),
	)
}