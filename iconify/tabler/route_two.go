package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouteTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 17l4 4m0-4l-4 4M17 3l4 4m0-4l-4 4m-3-2a2 2 0 0 0-2 2v10a2 2 0 0 1-2 2"/>`),
		g.Group(children),
	)
}