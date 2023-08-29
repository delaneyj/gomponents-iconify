package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipHorizontalTwoArrowDesignFlipReflect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.5 3.75l2.93 3.07a.27.27 0 0 1 0 .36L1.5 10.25m11-6.5L9.57 6.82a.27.27 0 0 0 0 .36l2.93 3.07M7 .5v1m0 2v1m0 2v1m0 2v1m0 2v1"/>`),
		g.Group(children),
	)
}