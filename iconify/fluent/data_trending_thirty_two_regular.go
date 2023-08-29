package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTrendingThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3a1 1 0 0 1 1 1v20.5A2.5 2.5 0 0 0 7.5 27H28a1 1 0 1 1 0 2H7.5A4.5 4.5 0 0 1 3 24.5V4a1 1 0 0 1 1-1Zm15 5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v7a1 1 0 1 1-2 0v-4.586l-7.293 7.293a1 1 0 0 1-1.414 0L14 14.414l-5.293 5.293a1 1 0 0 1-1.414-1.414l6-6a1 1 0 0 1 1.414 0L18 15.586L24.586 9H20a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}