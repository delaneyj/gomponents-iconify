package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderManSymbolGeometricGenderBoyPersonMaleHumanUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.5" cy="9.5" r="4"/><path d="M9 .5h4.5V5M7.33 6.67L13.5.5"/></g>`),
		g.Group(children),
	)
}