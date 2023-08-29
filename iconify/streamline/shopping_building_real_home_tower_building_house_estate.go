package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBuildingRealHomeTowerBuildingHouseEstate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 13.5h-8V4l4-3.5l4 3.5v9.5zm0 0h5v-7h-5m-4 7v-2M3 8.5h3m-3-3h3"/>`),
		g.Group(children),
	)
}