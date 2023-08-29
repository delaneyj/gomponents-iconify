package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineDashesFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M43.56 6.56a1.5 1.5 0 0 0-2.12-2.12l-2 2a1.5 1.5 0 0 0 2.12 2.12l2-2Zm-8 5.88a1.5 1.5 0 0 1 0 2.12l-3 3a1.5 1.5 0 0 1-2.12-2.12l3-3a1.5 1.5 0 0 1 2.12 0Zm-9 11.12a1.5 1.5 0 0 0-2.12-2.12l-3 3a1.5 1.5 0 0 0 2.12 2.12l3-3Zm-9 6.88a1.5 1.5 0 0 1 0 2.12l-3 3a1.5 1.5 0 0 1-2.12-2.12l3-3a1.5 1.5 0 0 1 2.12 0Zm-9 9a1.5 1.5 0 0 1 0 2.12l-2 2a1.5 1.5 0 0 1-2.12-2.12l2-2a1.5 1.5 0 0 1 2.12 0Z"/>`),
		g.Group(children),
	)
}