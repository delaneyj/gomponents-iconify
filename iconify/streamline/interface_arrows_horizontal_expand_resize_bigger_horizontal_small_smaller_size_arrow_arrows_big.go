package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsHorizontalExpandResizeBiggerHorizontalSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7H.5m3-3l-3 3l3 3m7-6l3 3l-3 3"/>`),
		g.Group(children),
	)
}