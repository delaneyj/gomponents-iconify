package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsCurvyLeftSnakeArrowSidewaysDiagramLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.75h-12a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h11a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H.5"/><path d="m3 8.25l-2.5 2.5l2.5 2.5"/></g>`),
		g.Group(children),
	)
}