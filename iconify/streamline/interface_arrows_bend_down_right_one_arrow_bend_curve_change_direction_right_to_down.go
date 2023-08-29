package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendDownRightOneArrowBendCurveChangeDirectionRightToDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 10.5l-3 3l-3-3"/><path d="M2.5.5h2a4 4 0 0 1 4 4v9"/></g>`),
		g.Group(children),
	)
}