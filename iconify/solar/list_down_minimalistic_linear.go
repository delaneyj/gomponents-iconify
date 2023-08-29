package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDownMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 6H3m17 5H3m7 5H3"/><path stroke-linejoin="round" d="m14 15l3.5 3l3.5-3"/></g>`),
		g.Group(children),
	)
}