package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeClockNineToFiveWorkHourTimeClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 12v1.5m-3.54-2.96L2.4 11.6M7 .5A6.5 6.5 0 0 0 .5 7H7l4.6 4.6A6.51 6.51 0 0 0 7 .5Z"/>`),
		g.Group(children),
	)
}