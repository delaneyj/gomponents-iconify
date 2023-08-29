package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelServingDomeHandPorterServiceRoomPlateHandBellhopPlatterGiveFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 5.5a5 5 0 0 1 10 0m-11 0h12m-13 3H4A2.5 2.5 0 0 1 6.5 11h0m-3 0h5a2 2 0 0 1 2 2h0a.5.5 0 0 1-.5.5H.5"/>`),
		g.Group(children),
	)
}