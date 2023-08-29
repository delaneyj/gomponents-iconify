package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderWomanSymbolGeometricGenderFemalePersonHumanUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="4" r="3.5"/><path d="M7 7.5v6M5 11h4"/></g>`),
		g.Group(children),
	)
}