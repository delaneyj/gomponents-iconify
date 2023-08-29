package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadBranches(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 4V1H0v3h1.7l7.7 9.5c1.3 1.6 3.1 2.5 5 2.5H16v-3h-1.5c-1 0-1.9-.5-2.7-1.4L10.5 10H16V7H8L5.6 4H16z"/>`),
		g.Group(children),
	)
}