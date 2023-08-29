package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDownLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M21 6H3m18 4H3m7 4H3m7 4H3" opacity=".5"/><path stroke-linejoin="round" d="m14 15l3.5 3l3.5-3"/></g>`),
		g.Group(children),
	)
}