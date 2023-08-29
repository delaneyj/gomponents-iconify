package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 3a1 1 0 0 1 1 1v5.25a.75.75 0 0 1-1.5 0V5.559L5.559 15.5H9.25a.75.75 0 0 1 0 1.5H4a1 1 0 0 1-1-1v-5.25a.75.75 0 0 1 1.5 0v3.689L14.439 4.5H10.75a.75.75 0 0 1 0-1.5H16Z"/>`),
		g.Group(children),
	)
}