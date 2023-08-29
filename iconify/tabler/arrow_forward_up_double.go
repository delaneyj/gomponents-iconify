package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForwardUpDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m11 14l4-4l-4-4m5 8l4-4l-4-4"/><path d="M15 10H8a4 4 0 1 0 0 8h1"/></g>`),
		g.Group(children),
	)
}