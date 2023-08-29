package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnRightDownFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 13a1.5 1.5 0 0 1 0-3h18a7.5 7.5 0 0 1 7.5 7.5v17.379l6.44-6.44a1.5 1.5 0 0 1 2.12 2.122l-9 9a1.5 1.5 0 0 1-2.12 0l-9-9a1.5 1.5 0 0 1 2.12-2.122L30 34.88V17.5a4.5 4.5 0 0 0-4.5-4.5h-18Z"/>`),
		g.Group(children),
	)
}