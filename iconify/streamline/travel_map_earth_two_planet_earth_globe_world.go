package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapEarthTwoPlanetEarthGlobeWorld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M12.48 3.5h-4a2 2 0 0 0 0 4a1 1 0 0 1 1 1V13M.58 8H3a2 2 0 0 1 2 2v3.19"/></g>`),
		g.Group(children),
	)
}