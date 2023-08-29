package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderStairsTwoStairsStaircaseLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5.5H5V5h4v4h4.5v4.5H.5V.5z"/>`),
		g.Group(children),
	)
}