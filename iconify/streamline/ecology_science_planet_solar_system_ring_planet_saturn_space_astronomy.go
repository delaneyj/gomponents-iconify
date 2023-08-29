package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EcologySciencePlanetSolarSystemRingPlanetSaturnSpaceAstronomy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.63 8.13C.85 10.49.05 12.53.76 13.24c1 1 4.6-1 8-4.44s5.44-7 4.44-8c-.64-.65-2.38 0-4.47 1.41"/><path d="M12.05 4.92A5 5 0 0 1 7.5 12a5.06 5.06 0 0 1-1.95-.39M3.5 10a5 5 0 0 1 7-7"/></g>`),
		g.Group(children),
	)
}