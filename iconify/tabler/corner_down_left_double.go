package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeftDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 5v6a3 3 0 0 1-3 3H9"/><path d="m13 10l-4 4l4 4m-5-8l-4 4l4 4"/></g>`),
		g.Group(children),
	)
}