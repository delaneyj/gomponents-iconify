package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHomeFiveDoorEntranceHomeHouseMapRoofRoundWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 8.5L7 2l6.5 6.5"/><path d="M2.5 6.5v7h9v-7m-4.5 7v-3"/><circle cx="7" cy="6.75" r="1.25"/></g>`),
		g.Group(children),
	)
}