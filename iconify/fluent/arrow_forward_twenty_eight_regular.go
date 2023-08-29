package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForwardTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.97 6.53a.75.75 0 0 1 1.06-1.06l5.75 5.75a.75.75 0 0 1 0 1.06l-5.75 5.75a.75.75 0 1 1-1.06-1.06l4.47-4.47h-8.69a9.25 9.25 0 0 0-9.25 9.25a.75.75 0 0 1-1.5 0C3 15.813 7.813 11 13.75 11h8.69l-4.47-4.47Z"/>`),
		g.Group(children),
	)
}