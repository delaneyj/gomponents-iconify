package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelTransportationCarTransportationTravelTaxiTransportCab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 10.5a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H11l-.77-2.32a1 1 0 0 0-1-.68H5.22a1 1 0 0 0-1 .68L3.5 6.5h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1"/><circle cx="3.5" cy="10.5" r="2"/><circle cx="10.5" cy="10.5" r="2"/><path d="M5.5 10.5h3M7 3.5v-2"/></g>`),
		g.Group(children),
	)
}