package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackPop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 9.5L4 11l8 4l8-4l-3-1.5M4 15l8 4l8-4m-8-4V4M9 7l3-3l3 3"/>`),
		g.Group(children),
	)
}