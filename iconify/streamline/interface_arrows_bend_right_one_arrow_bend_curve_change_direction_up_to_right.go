package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendRightOneArrowBendCurveChangeDirectionUpToRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11 2.75l2.5 2.5l-2.5 2.5"/><path d="M.5 11.25v-5a1 1 0 0 1 1-1h12"/></g>`),
		g.Group(children),
	)
}