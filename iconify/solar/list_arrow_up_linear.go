package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListArrowUpLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M21 6H3m18 4H3m8 4H3m8 4H3"/><path stroke-linejoin="round" d="M20 16.5L17.5 14m0 0L15 16.5m2.5-2.5v6"/></g>`),
		g.Group(children),
	)
}