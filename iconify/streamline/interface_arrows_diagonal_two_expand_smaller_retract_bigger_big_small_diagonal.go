package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsDiagonalTwoExpandSmallerRetractBiggerBigSmallDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5.5h-4v4m0-4l13 13m-4 0h4v-4"/>`),
		g.Group(children),
	)
}