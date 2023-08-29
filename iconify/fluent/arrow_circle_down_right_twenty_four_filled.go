package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm2.5-13.25a.75.75 0 0 1 1.5 0v6.5a.75.75 0 0 1-.75.75h-6.5a.75.75 0 0 1 0-1.5h4.74L8.222 9.283a.75.75 0 0 1 1.056-1.066L14.5 13.39V8.75Z"/>`),
		g.Group(children),
	)
}