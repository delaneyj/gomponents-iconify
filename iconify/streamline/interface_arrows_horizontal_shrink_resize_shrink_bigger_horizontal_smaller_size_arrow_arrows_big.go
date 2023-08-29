package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsHorizontalShrinkResizeShrinkBiggerHorizontalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 4.5L8.5 7L11 9.5m-8-5L5.5 7L3 9.5M.5.5v13m13-13v13"/>`),
		g.Group(children),
	)
}