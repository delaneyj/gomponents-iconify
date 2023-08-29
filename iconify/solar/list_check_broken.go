package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheckBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m14 16l2.1 2.5l3.9-5"/><path d="M10 14H3m7 4H3M3 6h10.5M20 6h-2.25M20 10H9.5M3 10h2.25"/></g>`),
		g.Group(children),
	)
}