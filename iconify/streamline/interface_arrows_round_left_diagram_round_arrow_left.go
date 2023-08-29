package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsRoundLeftDiagramRoundArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.75.5L2.25 3l2.5 2.5"/><path d="M1.75 8.25A5.25 5.25 0 1 0 7 3H2.25"/></g>`),
		g.Group(children),
	)
}