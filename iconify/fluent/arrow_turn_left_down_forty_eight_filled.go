package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnLeftDownFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M40.5 13a1.5 1.5 0 0 0 0-3h-18a7.5 7.5 0 0 0-7.5 7.5v17.379l-6.44-6.44a1.5 1.5 0 0 0-2.12 2.122l9 9a1.5 1.5 0 0 0 2.12 0l9-9a1.5 1.5 0 0 0-2.12-2.122L18 34.88V17.5a4.5 4.5 0 0 1 4.5-4.5h18Z"/>`),
		g.Group(children),
	)
}