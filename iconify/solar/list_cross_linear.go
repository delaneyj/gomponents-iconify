package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCrossLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m15 18.5l5-5m0 5l-5-5"/><path d="M21 6H3m18 4H3m8 4H3m8 4H3"/></g>`),
		g.Group(children),
	)
}