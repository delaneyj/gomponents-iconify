package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarWeekEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 8h14v11H5M6 1v2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14c1.11 0 2-.89 2-2V5c0-1.11-.89-2-2-2h-1V1h-2v2H8V1m7 16h2v-7h-2v7Z"/>`),
		g.Group(children),
	)
}