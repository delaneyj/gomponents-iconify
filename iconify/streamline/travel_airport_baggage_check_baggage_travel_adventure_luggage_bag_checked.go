package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelAirportBaggageCheckBaggageTravelAdventureLuggageBagChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9.5" x=".5" y="4" rx="2"/><path d="M4 13.5V4m6 9.5V4M4.5 4a2.5 2.5 0 0 1 5 0"/></g>`),
		g.Group(children),
	)
}