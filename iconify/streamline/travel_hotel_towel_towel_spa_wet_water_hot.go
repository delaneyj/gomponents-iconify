package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelTowelTowelSpaWetWaterHot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 10a3 3 0 0 0-3-3h-9a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h9a3 3 0 0 0 3-3Zm-13 0h9m-6-3a3 3 0 0 1 0-6h9a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-1.86m2.86-3h-9"/>`),
		g.Group(children),
	)
}