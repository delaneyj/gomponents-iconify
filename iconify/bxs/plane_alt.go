package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaneAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.414 13.778L2 15.192l4.949 2.121l2.122 4.95l1.414-1.414l-.707-3.536L13.091 14l3.61 7.704l1.339-1.339l-1.19-10.123l2.828-2.829a2 2 0 1 0-2.828-2.828l-2.903 2.903L3.824 6.297L2.559 7.563l7.644 3.67l-3.253 3.253l-3.536-.708z"/>`),
		g.Group(children),
	)
}