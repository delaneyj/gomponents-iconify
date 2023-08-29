package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelPoolLadderTwoPoolStairsSwimSwimmingWaterLadder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5H13a2 2 0 0 1-2-2a2 2 0 0 1-4 0a2 2 0 0 1-4 0a2 2 0 0 1-2 2H.5m2-11a2 2 0 0 1 4 0v7m3-9a2 2 0 0 1 2 2v7m-5-5h5m-5 3h5"/>`),
		g.Group(children),
	)
}