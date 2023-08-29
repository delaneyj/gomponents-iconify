package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapGlobeModelPlanetEarthGlobeWorld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.16" cy="5.75" r="3"/><path d="M6.16 11v2.5m-1.5 0h3M6.16.5a5.25 5.25 0 1 1-3.57 9.1"/></g>`),
		g.Group(children),
	)
}