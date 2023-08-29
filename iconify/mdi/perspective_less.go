package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerspectiveLess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.56 12l2.67 8H4.78l2.66-8h9.12M7 1L3 5l4 4V6h4V4H7V1m10 0v3h-4v2h4v3l4-4l-4-4m1 9H6L2 22h20l-4-12Z"/>`),
		g.Group(children),
	)
}