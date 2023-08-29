package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsTurnUpArrowBendCurveChangeDirectionReturnUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12.5 3.5l-3-3l-3 3"/><path d="M9.5.5v9a4 4 0 0 1-8 0v-5"/></g>`),
		g.Group(children),
	)
}