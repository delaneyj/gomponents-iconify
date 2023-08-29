package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelDumbellSportsWeightsDumbbellSportFitness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="4" height="6" x=".5" y="4" rx="1"/><rect width="4" height="6" x="9.5" y="4" rx="1"/><path d="M4.5 7h5"/></g>`),
		g.Group(children),
	)
}