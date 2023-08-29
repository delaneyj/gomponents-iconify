package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarPersonSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.367 14H4.5A2.5 2.5 0 0 1 2 11.5V6h7.264A2.989 2.989 0 0 0 8.5 8c0 .773.293 1.478.773 2.01A2.5 2.5 0 0 0 7 12.5c0 .512.118 1.027.367 1.5ZM14 5v-.5A2.5 2.5 0 0 0 11.5 2h-7A2.5 2.5 0 0 0 2 4.5V5h12Zm-.5 3a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm1.5 4.5c0 1.245-1 2.5-3.5 2.5S8 13.75 8 12.5A1.5 1.5 0 0 1 9.5 11h4a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}