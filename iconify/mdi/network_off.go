package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1 5.27l4 4V15a2 2 0 0 0 2 2h4v2h-1a1 1 0 0 0-1 1H2v2h7a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1h2.73l2 2L21 22.72L2.28 4L1 5.27M15 20a1 1 0 0 0-1-1h-1v-1.73L15.73 20H15m2.69-3.13L5.13 4.31C5.41 3.55 6.14 3 7 3h10a2 2 0 0 1 2 2v10c0 .86-.55 1.59-1.31 1.87M22 20v1.18L20.82 20H22Z"/>`),
		g.Group(children),
	)
}