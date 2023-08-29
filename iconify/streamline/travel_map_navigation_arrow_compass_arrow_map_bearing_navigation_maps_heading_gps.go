package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapNavigationArrowCompassArrowMapBearingNavigationMapsHeadingGps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 13.5l5-13l-13 5l6 2l2 6z"/>`),
		g.Group(children),
	)
}