package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarWeekOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3h1V1h2v2h8V1h2v2h1c1.11 0 2 .89 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2m0 16h14V9H5v10M5 7h14V5H5v2m12 4v2H7v-2h10"/>`),
		g.Group(children),
	)
}