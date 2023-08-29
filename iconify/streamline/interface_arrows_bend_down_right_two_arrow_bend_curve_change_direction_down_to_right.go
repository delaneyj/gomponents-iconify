package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendDownRightTwoArrowBendCurveChangeDirectionDownToRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 11.5l3-3l-3-3"/><path d="M.5 2.5v2a4 4 0 0 0 4 4h9"/></g>`),
		g.Group(children),
	)
}