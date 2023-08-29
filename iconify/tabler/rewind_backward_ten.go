package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindBackwardTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 9L4 6l3-3"/><path d="M15.997 17.918A6.002 6.002 0 0 0 15 6H4m2 8v6m3-4.5v3a1.5 1.5 0 0 0 3 0v-3a1.5 1.5 0 0 0-3 0z"/></g>`),
		g.Group(children),
	)
}