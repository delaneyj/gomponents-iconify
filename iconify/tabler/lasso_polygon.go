package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LassoPolygon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.028 13.252L3 10l2-7l7 5l8-3l1 9l-9 3l-5.144-1.255"/><path d="M3 15a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M5 17c0 1.42.316 2.805 1 4"/></g>`),
		g.Group(children),
	)
}