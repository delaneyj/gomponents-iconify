package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewWeekOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 18h2.5V6H13v12Zm-4.5 0H11V6H8.5v12ZM4 18h2.5V6H4v12Zm13.5 0H20V6h-2.5v12ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}