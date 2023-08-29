package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapFlagNavigationMapMapsFlagGpsLocationDestinationGoal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5v13m9-6h-9v-7h9L8.5 4l3 3.5z"/>`),
		g.Group(children),
	)
}