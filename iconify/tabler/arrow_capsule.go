package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCapsule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 15a6 6 0 1 1-12 0V9a6 6 0 1 1 12 0v2"/><path d="m15 8l3 3l3-3"/></g>`),
		g.Group(children),
	)
}