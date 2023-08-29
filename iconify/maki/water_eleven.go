package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 11C3.59 11 2 9 2 7s2.61-5.81 3.5-7C6.39 1.19 9 5 9 7s-1.59 4-3.5 4z" fill="currentColor"/>`),
		g.Group(children),
	)
}