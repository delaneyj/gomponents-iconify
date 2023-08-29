package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderLadderBusinessProductMetaphorLadder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5.5v13m7-13v13M3.5 3h7m-7 4h7m-7 4h7"/>`),
		g.Group(children),
	)
}