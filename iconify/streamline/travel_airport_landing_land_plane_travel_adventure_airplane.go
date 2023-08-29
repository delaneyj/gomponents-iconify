package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelAirportLandingLandPlaneTravelAdventureAirplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3 1.93l1.38.54a.46.46 0 0 1 .3.47l-.18 1.8l5.84 1.35A3.37 3.37 0 0 1 13 10.5a.91.91 0 0 1-1.19.5L8.36 9.65L8 9.52l-2 1.34a.52.52 0 0 1-.47.05L4 10.3a.5.5 0 0 1-.14-.85l1.28-1.06l-.23-.09L1.54 7A.9.9 0 0 1 1 5.83l1.44-3.66A.44.44 0 0 1 3 1.93ZM.5 13.5h13"/>`),
		g.Group(children),
	)
}