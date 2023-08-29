package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCrossMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 6H3m8 5H3m8 5H3" opacity=".5"/><path stroke-linejoin="round" d="m15 16l5-5m0 5l-5-5"/></g>`),
		g.Group(children),
	)
}