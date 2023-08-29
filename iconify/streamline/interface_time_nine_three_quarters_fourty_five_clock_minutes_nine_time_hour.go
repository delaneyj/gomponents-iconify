package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeNineThreeQuartersFourtyFiveClockMinutesNineTimeHour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7A6.5 6.5 0 1 0 7 .5V7Zm2.96-3.54L2.4 2.4"/>`),
		g.Group(children),
	)
}