package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneCircleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm-4-7.566l-.861-.319a1.75 1.75 0 0 1-1.143-1.641v-.95A1.75 1.75 0 0 1 7.14 9.881l7.501-2.775a1.75 1.75 0 0 1 2.357 1.641v6.5a1.75 1.75 0 0 1-2.357 1.642l-1.434-.531A2.626 2.626 0 0 1 8 15.875v-1.441Zm3.75 1.387v.054a1.125 1.125 0 0 1-2.25 0v-.887l2.25.833Z"/>`),
		g.Group(children),
	)
}