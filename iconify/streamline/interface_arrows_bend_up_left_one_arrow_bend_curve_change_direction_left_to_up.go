package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendUpLeftOneArrowBendCurveChangeDirectionLeftToUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.75 3L5.25.5L7.75 3"/><path d="M11.25 13.5h-5a1 1 0 0 1-1-1V.5"/></g>`),
		g.Group(children),
	)
}