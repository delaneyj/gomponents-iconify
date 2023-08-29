package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewWeekSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.75 19V5h3.375v14H12.75Zm-4.875 0V5h3.375v14H7.875ZM3 19V5h3.375v14H3Zm14.625 0V5H21v14h-3.375Z"/>`),
		g.Group(children),
	)
}