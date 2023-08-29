package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHomeTwoDoorEntranceHomeHouseMapRoofRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 9L7 2.5L13.5 9"/><path d="M2.5 7v6.5h9V7M7 13.5v-4"/></g>`),
		g.Group(children),
	)
}