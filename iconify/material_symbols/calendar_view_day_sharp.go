package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewDaySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h18v2H3Zm0-4V7h18v10H3ZM3 5V3h18v2H3Z"/>`),
		g.Group(children),
	)
}