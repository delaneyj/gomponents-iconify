package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m7 8l-4 4l4 4"/><path d="m10.5 18l3-12"/><path stroke-linejoin="round" d="m17 8l4 4l-4 4"/></g>`),
		g.Group(children),
	)
}