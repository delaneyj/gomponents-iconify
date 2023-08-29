package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWafinderSinkWashCleanToiletBathroomWater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5h13v2a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3v-2h0Zm7-6a2 2 0 0 0-4 0v6m4-4v1"/>`),
		g.Group(children),
	)
}