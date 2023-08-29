package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelBathTubBatheBathBathroomTowelWaterTub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 4.5h-5a1 1 0 0 0-1 1v1a4 4 0 0 0 4 4h5a4 4 0 0 0 4-4v-1a1 1 0 0 0-1-1h-2"/><path d="M6.5 3.5h4v4h-4z"/></g>`),
		g.Group(children),
	)
}