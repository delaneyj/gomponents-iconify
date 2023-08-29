package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapLocationCompassOneArrowCompassLocationGpsMapMapsPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m7.5 10.5l2-6l-6 2L6 8l1.5 2.5z"/></g>`),
		g.Group(children),
	)
}