package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingChurch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 21h18m-11 0v-4a2 2 0 0 1 4 0v4M10 5h4m-2-2v5"/><path d="M6 21v-7m-2 2l8-8l8 8m-2-2v7"/></g>`),
		g.Group(children),
	)
}