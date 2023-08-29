package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IterationCw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10c0-4.4 3.6-8 8-8s8 3.6 8 8s-3.6 8-8 8H4"/><path d="m8 22l-4-4l4-4"/></g>`),
		g.Group(children),
	)
}