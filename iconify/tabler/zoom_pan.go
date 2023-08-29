package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomPan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m8 5l-2.5-2.5M10 5l2-2l2 2m5 5l2 2l-2 2M5 10l-2 2l2 2m5 5l2 2l2-2"/>`),
		g.Group(children),
	)
}