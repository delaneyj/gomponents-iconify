package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCrossMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M3 6h10.5M20 6h-2.25M11 16H3"/><path stroke-linejoin="round" d="m15 16l5-5m0 5l-5-5"/><path d="M11 11H7m-4 0h1.2"/></g>`),
		g.Group(children),
	)
}