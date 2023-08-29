package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelTransportationBusTransportationTravelBusTransitTransportMotorcoachPublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="11.5" x="1.5" y=".5" rx="1"/><path d="M3.5 12v1.5m7-1.5v1.5M1.5 7h11"/><circle cx="4" cy="9.5" r=".5"/><circle cx="10" cy="9.5" r=".5"/></g>`),
		g.Group(children),
	)
}