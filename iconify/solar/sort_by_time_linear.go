package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortByTimeLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M10 7H2m6 5H2m8 5H2"/><circle cx="17" cy="12" r="5"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 10v1.846L18 13"/></g>`),
		g.Group(children),
	)
}