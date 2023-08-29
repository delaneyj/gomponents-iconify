package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsExpandDiagonalTwoExpandResizeBiggerDiagonalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 5L.5.5m4 0h-4v4M9 9l4.5 4.5m-4 0h4v-4M10 4l-6 6"/>`),
		g.Group(children),
	)
}