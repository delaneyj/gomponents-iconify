package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 19.966V4.034h-6.654zM3 4.034v15.932L9.658 4.034zM9.092 16.76h3.104l1.268 3.205h2.778L12.003 9.904z"/>`),
		g.Group(children),
	)
}