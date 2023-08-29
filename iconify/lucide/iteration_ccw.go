package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IterationCcw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10c0-4.4-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8h8"/><path d="m16 14l4 4l-4 4"/></g>`),
		g.Group(children),
	)
}