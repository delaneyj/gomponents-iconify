package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendDownLeftThreeArrowBendCurveChangeDirectionLeftToDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5 10.5l3 3l3-3"/><path d="M11.5.5h-2a4 4 0 0 0-4 4v9"/></g>`),
		g.Group(children),
	)
}