package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderLifebuoyWaterLifeRingWheelyLifebeltKisbee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 1.5L9.09 4.91M1.5 1.5l3.41 3.41M1.5 12.5l3.41-3.41m7.59 3.41L9.09 9.09"/><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="7" r="2.95"/></g>`),
		g.Group(children),
	)
}