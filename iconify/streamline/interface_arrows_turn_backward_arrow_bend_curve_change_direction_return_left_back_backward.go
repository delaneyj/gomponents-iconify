package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsTurnBackwardArrowBendCurveChangeDirectionReturnLeftBackBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 1.5l-3 3l3 3"/><path d="M.5 4.5h9a4 4 0 0 1 0 8h-5"/></g>`),
		g.Group(children),
	)
}