package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUpMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 11H3m8 5H3"/><path stroke-linejoin="round" d="m14 18l3.5-3l3.5 3"/><path d="M3 6h10.5M20 6h-2.25"/></g>`),
		g.Group(children),
	)
}