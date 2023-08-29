package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartRate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31.878 37c-2.825 2.074-5.654 3.587-7.878 4.326C17 39 4 29 4 18C4 11.925 8.925 7 15 7c3.72 0 7.01 1.847 9 4.674A10.987 10.987 0 0 1 33 7c6.075 0 11 4.925 11 11c0 1.747-.328 3.468-.907 5.137"/><path d="M27 29h4l3-4l3 8l2.962-4H44"/></g>`),
		g.Group(children),
	)
}