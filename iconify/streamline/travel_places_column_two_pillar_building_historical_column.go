package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesColumnTwoPillarBuildingHistoricalColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.75" cy="2.75" r="2.25"/><circle cx="11.25" cy="2.75" r="2.25"/><path d="M3 .5h8m-6.84 4h5.68M3 13.5V4.99m2.5 8.51v-6m3 6v-6m2.5 6V4.99"/></g>`),
		g.Group(children),
	)
}