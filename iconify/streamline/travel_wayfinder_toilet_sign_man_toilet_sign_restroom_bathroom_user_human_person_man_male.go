package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderToiletSignManToiletSignRestroomBathroomUserHumanPersonManMale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M7 13.5c-3 0-3-7-3-7h6s0 7-3 7Z"/></g>`),
		g.Group(children),
	)
}