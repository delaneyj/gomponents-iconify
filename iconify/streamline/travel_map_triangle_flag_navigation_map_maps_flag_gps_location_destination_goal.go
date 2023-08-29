package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapTriangleFlagNavigationMapMapsFlagGpsLocationDestinationGoal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5v13m0-13l9 4.5l-9 4.5"/>`),
		g.Group(children),
	)
}