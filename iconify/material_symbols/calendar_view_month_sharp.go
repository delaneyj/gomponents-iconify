package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewMonthSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11.5V5h5.325v6.5H3Zm6.325 0V5h5.35v6.5h-5.35Zm6.35 0V5H21v6.5h-5.325ZM3 19v-6.5h5.325V19H3Zm6.325 0v-6.5h5.35V19h-5.35Zm6.35 0v-6.5H21V19h-5.325Z"/>`),
		g.Group(children),
	)
}