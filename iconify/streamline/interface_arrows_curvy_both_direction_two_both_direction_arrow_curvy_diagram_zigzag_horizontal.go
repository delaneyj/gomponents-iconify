package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsCurvyBothDirectionTwoBothDirectionArrowCurvyDiagramZigzagHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 .5L.5 3h12a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-11a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h12L11 13.5"/>`),
		g.Group(children),
	)
}