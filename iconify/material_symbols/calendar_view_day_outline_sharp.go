package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewDayOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h18v2H3Zm0-4V7h18v10H3Zm2-2h14V9H5v6ZM3 5V3h18v2H3Zm2 10V9v6Z"/>`),
		g.Group(children),
	)
}