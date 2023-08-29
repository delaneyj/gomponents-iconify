package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnLeftUpFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M40.5 35a1.5 1.5 0 0 1 0 3h-18a7.5 7.5 0 0 1-7.5-7.5V13.121l-6.44 6.44a1.5 1.5 0 0 1-2.12-2.122l9-9a1.5 1.5 0 0 1 2.12 0l9 9a1.5 1.5 0 0 1-2.12 2.122L18 13.12V30.5a4.5 4.5 0 0 0 4.5 4.5h18Z"/>`),
		g.Group(children),
	)
}