package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsExpandDiagonalOneExpandResizeBiggerDiagonalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 5L13.5.5m-4 0h4v4M5 9L.5 13.5m4 0h-4v-4M4 4l6 6"/>`),
		g.Group(children),
	)
}