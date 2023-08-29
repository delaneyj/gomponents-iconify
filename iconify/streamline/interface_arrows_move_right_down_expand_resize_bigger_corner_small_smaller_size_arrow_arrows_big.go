package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsMoveRightDownExpandResizeBiggerCornerSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5.5l2 2l-2 2"/><path d="M2.5 13.5v-10a1 1 0 0 1 1-1h10"/><path d="m4.5 11.5l-2 2l-2-2"/></g>`),
		g.Group(children),
	)
}