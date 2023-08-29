package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesStreetSignCrossroadStreetSignMetaphorDirections(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 4.5H2.5l-2-2l2-2H7m0 7h4.5l2-2l-2-2H7m0 7H2.5l-2-2l2-2H7m0-6v13"/>`),
		g.Group(children),
	)
}