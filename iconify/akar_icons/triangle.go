package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M10.293 4.793c.78-1.277 2.634-1.277 3.414 0l7.433 12.164C21.955 18.29 20.996 20 19.434 20H4.566c-1.562 0-2.52-1.71-1.706-3.043l7.433-12.164Z"/>`),
		g.Group(children),
	)
}