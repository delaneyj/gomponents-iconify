package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderDisabilityPersonAccessWheelchairAccomodationHumanDisabilityDisabledUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="8" cy="2.5" r="2"/><path d="M8 4.5v4h3a1 1 0 0 1 1 1v4m-6.5-7a3.5 3.5 0 1 0 3.22 4.88"/></g>`),
		g.Group(children),
	)
}