package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendUpLeftTwoArrowBendCurveChangeDirectionLeftToUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5 3.5l3-3l3 3"/><path d="M11.5 13.5h-2a4 4 0 0 1-4-4v-9"/></g>`),
		g.Group(children),
	)
}