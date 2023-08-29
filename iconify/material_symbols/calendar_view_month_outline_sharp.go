package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewMonthOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 11h4V6H4v5Zm6 0h4V6h-4v5Zm6 0h4V6h-4v5ZM4 18h4v-5H4v5Zm6 0h4v-5h-4v5Zm6 0h4v-5h-4v5ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}