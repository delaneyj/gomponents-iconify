package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleLeftTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2c6.627 0 12 5.373 12 12s-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2Zm2.78 7.03a.75.75 0 0 0-1.06-1.06l-5.5 5.5a.75.75 0 0 0 0 1.06l5.5 5.5a.75.75 0 1 0 1.06-1.06L11.81 14l4.97-4.97Z"/>`),
		g.Group(children),
	)
}