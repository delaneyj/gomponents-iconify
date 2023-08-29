package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesSnorkleDivingScubaOutdoorRecreationOceanMaskWaterSeaSnorkle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 10.5a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V.5"/><path d="M10.5 5.5a2 2 0 0 0-2-2h-6a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h1l1.17-1.75a1 1 0 0 1 1.66 0L7.5 9.5h1a2 2 0 0 0 2-2Z"/></g>`),
		g.Group(children),
	)
}