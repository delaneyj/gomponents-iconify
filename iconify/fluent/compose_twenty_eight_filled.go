package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComposeTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.707 3.707a1 1 0 0 0-1.414-1.414l-13 13L11 17l1.707-.293l13-13ZM6.5 3A3.5 3.5 0 0 0 3 6.5v15A3.5 3.5 0 0 0 6.5 25h15a3.5 3.5 0 0 0 3.5-3.5v-10a1 1 0 1 0-2 0v10a1.5 1.5 0 0 1-1.5 1.5h-15A1.5 1.5 0 0 1 5 21.5v-15A1.5 1.5 0 0 1 6.5 5h10a1 1 0 1 0 0-2h-10Z"/>`),
		g.Group(children),
	)
}