package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 4H0v3h3.2L7 10.6c1.6 1.5 3.6 2.4 5.8 2.4H16v-3h-3.2c-1.4 0-2.7-.5-3.7-1.5L7.5 7H16V4z"/>`),
		g.Group(children),
	)
}