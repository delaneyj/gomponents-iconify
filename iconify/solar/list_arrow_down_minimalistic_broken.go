package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListArrowDownMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M11 16H3m9-5H3"/><path stroke-linejoin="round" d="m15 14.5l2.5 2.5m0 0l2.5-2.5M17.5 17V9"/><path d="M3 6h10.5M20 6h-2.25"/></g>`),
		g.Group(children),
	)
}