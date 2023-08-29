package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesColumnOnePillarBuildingHistoricalColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 13.25v-7.5m2.5 7.5v-5.5m3 5.5v-5.5m2.5 5.5v-7.5m0-5H3a2.5 2.5 0 1 0 2.29 3.5h3.42A2.5 2.5 0 1 0 11 .75Z"/>`),
		g.Group(children),
	)
}