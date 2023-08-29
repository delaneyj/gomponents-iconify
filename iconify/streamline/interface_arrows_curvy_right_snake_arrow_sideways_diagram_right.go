package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsCurvyRightSnakeArrowSidewaysDiagramRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.75h12a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-11a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h12"/><path d="m11 8.25l2.5 2.5l-2.5 2.5"/></g>`),
		g.Group(children),
	)
}