package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartCircles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 9.5a5.5 5.5 0 1 0 11 0a5.5 5.5 0 1 0-11 0"/><path d="M9 14.5a5.5 5.5 0 1 0 11 0a5.5 5.5 0 1 0-11 0"/></g>`),
		g.Group(children),
	)
}