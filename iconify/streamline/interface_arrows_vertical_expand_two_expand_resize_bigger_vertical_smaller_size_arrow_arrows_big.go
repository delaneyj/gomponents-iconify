package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsVerticalExpandTwoExpandResizeBiggerVerticalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 11.5v-9M9.5 5L7 2.5L4.5 5m5 4L7 11.5L4.5 9m9-8.5H.5m13 13H.5"/>`),
		g.Group(children),
	)
}