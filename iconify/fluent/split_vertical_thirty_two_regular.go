package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 29a1 1 0 1 0 2 0V3a1 1 0 1 0-2 0v26Zm11.5-5h-8v2h8a3.5 3.5 0 0 0 3.5-3.5v-13A3.5 3.5 0 0 0 26.5 6h-8v2h8A1.5 1.5 0 0 1 28 9.5v13a1.5 1.5 0 0 1-1.5 1.5Zm-13-18h-8A3.5 3.5 0 0 0 2 9.5v13A3.5 3.5 0 0 0 5.5 26h8v-2h-8A1.5 1.5 0 0 1 4 22.5v-13A1.5 1.5 0 0 1 5.5 8h8V6Z"/>`),
		g.Group(children),
	)
}