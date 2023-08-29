package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterclockwiseSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 8a4.5 4.5 0 0 0-7.329-3.5H6.25a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 3 5.25v-2.5a.75.75 0 0 1 1.5 0v.376A6 6 0 1 1 2 8a.75.75 0 0 1 1.5 0a4.5 4.5 0 1 0 9 0Z"/>`),
		g.Group(children),
	)
}