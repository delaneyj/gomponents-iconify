package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsTriangleLoopDiagramTriangleLoopArrowArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 13.5L4 11h8.5m1-7l-1 3.5L8 .5m-6 3h3.5l-5 7"/>`),
		g.Group(children),
	)
}