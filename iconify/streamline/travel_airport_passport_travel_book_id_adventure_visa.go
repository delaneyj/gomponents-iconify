package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelAirportPassportTravelBookIdAdventureVisa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><circle cx="7" cy="6" r="3"/><path d="M4 6h6M7 9V3"/></g>`),
		g.Group(children),
	)
}