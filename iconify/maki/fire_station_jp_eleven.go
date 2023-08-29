package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireStationJpEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.375 5.887v2.738a.875.875 0 0 1-1.75 0V5.887A3.75 3.75 0 0 1 1.75 2.25a.75.75 0 0 1 1.5 0a2.25 2.25 0 0 0 4.5 0a.75.75 0 0 1 1.5 0a3.75 3.75 0 0 1-2.875 3.637z" fill="currentColor"/>`),
		g.Group(children),
	)
}