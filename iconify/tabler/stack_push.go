package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackPush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 10l-2 1l8 4l8-4l-2-1M4 15l8 4l8-4M12 4v7"/><path d="m15 8l-3 3l-3-3"/></g>`),
		g.Group(children),
	)
}