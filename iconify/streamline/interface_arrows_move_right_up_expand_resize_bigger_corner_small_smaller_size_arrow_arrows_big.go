package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsMoveRightUpExpandResizeBiggerCornerSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m.5 2.5l2-2l2 2"/><path d="M13.5 11.5h-10a1 1 0 0 1-1-1V.5"/><path d="m11.5 9.5l2 2l-2 2"/></g>`),
		g.Group(children),
	)
}