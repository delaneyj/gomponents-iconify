package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.846 9.868l4.08.972A4 4 0 0 1 21 14.73V17m-3 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h2"/><path d="M8 18v-1a4 4 0 0 0-4-4H3m7-1l.663-1.327M3 3l18 18"/></g>`),
		g.Group(children),
	)
}