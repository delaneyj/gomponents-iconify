package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeThreeClockFifteenThreeQuarterMinutesTimeHour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7A6.5 6.5 0 0 0 7 .5V7ZM2 7H.5m2.96-3.54L2.4 2.4m1.06 8.14L2.4 11.6M7 12v1.5m3.54-2.96l1.06 1.06"/>`),
		g.Group(children),
	)
}