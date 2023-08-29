package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOneQuarterTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11 7.355l-.99 2.006l-6.017.874A1.447 1.447 0 0 0 2.75 11.68c0 .365.138.735.44 1.03l4.354 4.244l-1.028 5.992a1.436 1.436 0 0 0 .66 1.475c.41.257.946.316 1.444.054L11 23.223V7.355Z"/>`),
		g.Group(children),
	)
}