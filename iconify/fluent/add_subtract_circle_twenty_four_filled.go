package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSubtractCircleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.75 14.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5ZM12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm6.517-15.457A8.5 8.5 0 0 1 6.543 18.517L18.517 6.543ZM6 8.75A.75.75 0 0 1 6.75 8H8V6.75a.75.75 0 0 1 1.5 0V8h1.25a.75.75 0 0 1 0 1.5H9.5v1.25a.75.75 0 0 1-1.5 0V9.5H6.75A.75.75 0 0 1 6 8.75Z"/>`),
		g.Group(children),
	)
}