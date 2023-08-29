package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderLiftArrowUpHumanDownPersonUserLiftElevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="2.5" r="2"/><path d="M1.5 6.5h2a1 1 0 0 1 1 1v3h0h-4h0v-3a1 1 0 0 1 1-1Zm-1 4v3m4-3v3m4-8L11 3l2.5 2.5m-5 4L11 12l2.5-2.5"/></g>`),
		g.Group(children),
	)
}