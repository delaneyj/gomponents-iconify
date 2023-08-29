package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOffFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.31 6H4a2 2 0 0 0-2 2v16a2.85 2.85 0 0 0 0 .29zm9.15 2.11a1 1 0 0 0-1 .08L23 12.06v-1.62l7-7L28.56 2L2 28.56L3.44 30l4-4H21a2 2 0 0 0 2-2v-4.06l5.42 3.87A1 1 0 0 0 30 23V9a1 1 0 0 0-.54-.89z"/>`),
		g.Group(children),
	)
}