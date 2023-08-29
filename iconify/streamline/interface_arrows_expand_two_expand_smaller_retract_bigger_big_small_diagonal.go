package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsExpandTwoExpandSmallerRetractBiggerBigSmallDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 5.5l-5-5m4 0h-4v4m8 4l5 5m-4 0h4v-4"/>`),
		g.Group(children),
	)
}