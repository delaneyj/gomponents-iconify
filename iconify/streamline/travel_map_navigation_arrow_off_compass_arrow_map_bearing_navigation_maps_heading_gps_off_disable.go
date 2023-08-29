package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapNavigationArrowOffCompassArrowMapBearingNavigationMapsHeadingGpsOffDisable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.28 3.28L.5 5.5l6 2l2 6l2.22-5.78m1.08-2.8L13.5.5L8.91 2.27M3.5.5l10 10"/>`),
		g.Group(children),
	)
}