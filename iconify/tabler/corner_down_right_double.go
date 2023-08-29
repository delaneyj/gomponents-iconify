package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownRightDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 5v6a3 3 0 0 0 3 3h7"/><path d="m10 10l4 4l-4 4m5-8l4 4l-4 4"/></g>`),
		g.Group(children),
	)
}