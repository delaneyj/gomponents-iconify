package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 13v-1c0-.2 0-4.1-2.8-5.4C9 5.6 9 3.1 9 3V0H7v3c0 .1 0 2.6-2.2 3.6C2 7.9 2 11.8 2 12v1H0l3 3l3-3H4v-1s0-2.8 1.7-3.6c1.1-.5 1.8-1.3 2.3-2c.5.8 1.2 1.5 2.3 2C12 9.2 12 12 12 12v1h-2l3 3l3-3h-2z"/>`),
		g.Group(children),
	)
}