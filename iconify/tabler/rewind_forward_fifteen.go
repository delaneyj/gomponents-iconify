package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindForwardFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 9l3-3l-3-3"/><path d="M9 18A6 6 0 1 1 9 6h11m-4 14h2a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1h-2v-3h3m-6 0v6"/></g>`),
		g.Group(children),
	)
}