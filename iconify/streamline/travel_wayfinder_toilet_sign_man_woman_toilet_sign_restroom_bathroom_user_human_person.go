package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderToiletSignManWomanToiletSignRestroomBathroomUserHumanPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="2.5" r="2"/><path d="M11 6.5c-2.5 0-2.5 7-2.5 7h5s0-7-2.5-7Z"/><circle cx="3" cy="2.5" r="2"/><path d="M3 13.5c-2.5 0-2.5-7-2.5-7h5s0 7-2.5 7Z"/></g>`),
		g.Group(children),
	)
}