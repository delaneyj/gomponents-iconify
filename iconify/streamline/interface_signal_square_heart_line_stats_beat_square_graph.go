package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSignalSquareHeartLineStatsBeatSquareGraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="M2.5 7.02h2L6 4.51l1.5 5.5l2-2.99h2"/></g>`),
		g.Group(children),
	)
}