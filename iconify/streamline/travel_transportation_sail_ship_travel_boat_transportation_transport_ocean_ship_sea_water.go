package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelTransportationSailShipTravelBoatTransportationTransportOceanShipSeaWater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 9.5H.5l1 2.7a2 2 0 0 0 1.88 1.3h7.22a2 2 0 0 0 1.88-1.3Zm-7 0v-9m0 0l-5 6h5m2-4l3.5 4H8.5"/>`),
		g.Group(children),
	)
}