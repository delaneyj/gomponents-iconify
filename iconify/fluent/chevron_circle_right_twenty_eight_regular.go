package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleRightTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 24.5C8.201 24.5 3.5 19.799 3.5 14S8.201 3.5 14 3.5S24.5 8.201 24.5 14S19.799 24.5 14 24.5ZM2 14c0 6.627 5.373 12 12 12s12-5.373 12-12S20.627 2 14 2S2 7.373 2 14Zm9.22 4.97a.75.75 0 1 0 1.06 1.06l5.5-5.5a.75.75 0 0 0 0-1.06l-5.5-5.5a.75.75 0 1 0-1.06 1.06L16.19 14l-4.97 4.97Z"/>`),
		g.Group(children),
	)
}