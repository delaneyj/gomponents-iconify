package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOffFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26.74 19.56l-2.52-1a2 2 0 0 0-2.15.44L20 21.06a9.93 9.93 0 0 1-5.35-2.29L30 3.41L28.59 2L2 28.59L3.41 30l7.93-7.92c3.24 3.12 7.89 5.5 14.55 5.92A2 2 0 0 0 28 26v-4.59a2 2 0 0 0-1.26-1.85zM8.15 18.19l3.52-3.52a11.68 11.68 0 0 1-.82-2.67l2.07-2.07a2 2 0 0 0 .44-2.15l-1-2.52A2 2 0 0 0 10.5 4H6a2 2 0 0 0-2 2.22a29 29 0 0 0 4.15 11.97z"/>`),
		g.Group(children),
	)
}