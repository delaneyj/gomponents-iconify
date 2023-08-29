package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRemoveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3h-1V1h-2v2H8V1H6v2H5c-1.1 0-2 .9-2 2v14a2 2 0 0 0 2 2h14c1.11 0 2-.89 2-2V5a2 2 0 0 0-2-2m0 16H5V9h14v10M5 7V5h14v2H5m3.23 9.41l1.06 1.06l2.44-2.44l2.44 2.44l1.06-1.06l-2.44-2.44l2.44-2.44l-1.06-1.06l-2.44 2.44l-2.44-2.44l-1.06 1.06l2.44 2.44l-2.44 2.44Z"/>`),
		g.Group(children),
	)
}