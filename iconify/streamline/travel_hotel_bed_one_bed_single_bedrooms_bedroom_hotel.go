package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelBedOneBedSingleBedroomsBedroomHotel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><path d="M9.5.5v3h-5v-3M1.5 6h11"/></g>`),
		g.Group(children),
	)
}