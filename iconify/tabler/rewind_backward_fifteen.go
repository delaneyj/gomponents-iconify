package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindBackwardFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 20h2a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1H8v-3h3m4 4a6 6 0 1 0 0-12H4m1 8v6"/><path d="M7 9L4 6l3-3"/></g>`),
		g.Group(children),
	)
}