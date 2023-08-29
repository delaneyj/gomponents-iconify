package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsVerticalLeftRightExpandResizeBiggerHorizontalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.48 3.87L.61 6.74a.36.36 0 0 0 0 .52l2.87 2.87m7.04-6.26l2.87 2.87a.36.36 0 0 1 0 .52l-2.87 2.87M7 13.5V.5"/>`),
		g.Group(children),
	)
}