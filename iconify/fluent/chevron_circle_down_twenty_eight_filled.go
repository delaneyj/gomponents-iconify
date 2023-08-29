package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleDownTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 14C2 7.373 7.373 2 14 2s12 5.373 12 12s-5.373 12-12 12S2 20.627 2 14Zm7.03-2.78a.75.75 0 0 0-1.06 1.06l5.5 5.5a.75.75 0 0 0 1.06 0l5.5-5.5a.75.75 0 1 0-1.06-1.06L14 16.19l-4.97-4.97Z"/>`),
		g.Group(children),
	)
}