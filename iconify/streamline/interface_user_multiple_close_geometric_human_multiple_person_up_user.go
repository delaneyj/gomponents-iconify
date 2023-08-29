package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserMultipleCloseGeometricHumanMultiplePersonUpUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="3.75" r="2.25"/><path d="M9.5 13.5h-9v-1a4.5 4.5 0 0 1 9 0ZM9 1.5A2.25 2.25 0 0 1 9 6m1.6 2.19a4.5 4.5 0 0 1 2.9 4.2v1.11H12"/></g>`),
		g.Group(children),
	)
}