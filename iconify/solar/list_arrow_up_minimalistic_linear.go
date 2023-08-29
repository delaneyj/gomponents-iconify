package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListArrowUpMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 6H3m8 5H3m9 5H3"/><path stroke-linejoin="round" d="M15 11.5L17.5 9m0 0l2.5 2.5M17.5 9v8"/></g>`),
		g.Group(children),
	)
}