package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadblockEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 0a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11zM2 4.5h7v2H2v-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}