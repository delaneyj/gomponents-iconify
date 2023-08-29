package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.175 11.178a2 2 0 1 0 2.653 2.634M14.5 10.5l1-1"/><path d="M8.621 4.612a9 9 0 0 1 11.721 11.72m-1.516 2.488A9.008 9.008 0 0 1 17.6 20H6.4a9 9 0 0 1-.268-13.87M3 3l18 18"/></g>`),
		g.Group(children),
	)
}