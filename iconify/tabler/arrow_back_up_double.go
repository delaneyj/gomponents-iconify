package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBackUpDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13 14l-4-4l4-4m-5 8l-4-4l4-4"/><path d="M9 10h7a4 4 0 1 1 0 8h-1"/></g>`),
		g.Group(children),
	)
}