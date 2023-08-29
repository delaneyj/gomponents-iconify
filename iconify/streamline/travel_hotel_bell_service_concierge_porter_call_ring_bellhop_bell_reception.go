package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelBellServiceConciergePorterCallRingBellhopBellReception(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 10.25a6.5 6.5 0 0 1 13 0Zm0 2.5h13m-6.5-9v-2.5m-1.5 0h3"/>`),
		g.Group(children),
	)
}