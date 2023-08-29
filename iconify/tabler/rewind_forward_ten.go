package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindForwardTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 9l3-3l-3-3"/><path d="M8 17.918A5.997 5.997 0 0 1 3 12a6 6 0 0 1 6-6h11m-8 8v6m3-4.5v3a1.5 1.5 0 0 0 3 0v-3a1.5 1.5 0 0 0-3 0z"/></g>`),
		g.Group(children),
	)
}