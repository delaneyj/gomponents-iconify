package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsTurnForwardArrowBendCurveChangeDirectionReturnRightNextForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 1.5l3 3l-3 3"/><path d="M13.5 4.5h-9a4 4 0 0 0 0 8h5"/></g>`),
		g.Group(children),
	)
}