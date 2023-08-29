package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowClockwiseSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 8a4.5 4.5 0 0 1 7.329-3.5H9.75a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 .75-.75v-2.5a.75.75 0 0 0-1.5 0v.376A6 6 0 1 0 14 8a.75.75 0 0 0-1.5 0a4.5 4.5 0 1 1-9 0Z"/>`),
		g.Group(children),
	)
}