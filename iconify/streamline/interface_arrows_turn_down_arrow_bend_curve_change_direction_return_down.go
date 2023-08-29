package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsTurnDownArrowBendCurveChangeDirectionReturnDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.5 10.5l3 3l3-3"/><path d="M4.5 13.5v-9a4 4 0 0 1 8 0v5"/></g>`),
		g.Group(children),
	)
}