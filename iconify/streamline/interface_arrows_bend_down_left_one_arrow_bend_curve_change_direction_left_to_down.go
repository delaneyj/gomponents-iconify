package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendDownLeftOneArrowBendCurveChangeDirectionLeftToDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.75 11l2.5 2.5l2.5-2.5"/><path d="M11.25.5h-5a1 1 0 0 0-1 1v12"/></g>`),
		g.Group(children),
	)
}