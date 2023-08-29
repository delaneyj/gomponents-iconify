package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundAltArrowDownLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="m15 10.5l-3 3l-3-3"/></g>`),
		g.Group(children),
	)
}