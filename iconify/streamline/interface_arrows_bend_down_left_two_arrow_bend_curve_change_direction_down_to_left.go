package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendDownLeftTwoArrowBendCurveChangeDirectionDownToLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 11.5l-3-3l3-3"/><path d="M13.5 2.5v2a4 4 0 0 1-4 4h-9"/></g>`),
		g.Group(children),
	)
}